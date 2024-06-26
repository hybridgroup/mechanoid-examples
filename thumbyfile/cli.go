package main

import (
	"bytes"
	"runtime"
	"time"

	"encoding/binary"
	"encoding/hex"
	"strings"

	"github.com/hybridgroup/mechanoid/convert"
	"github.com/hybridgroup/mechanoid/engine"
)

const consoleBufLen = 64

const (
	StateInput = iota
	StateEscape
	StateEscBrc
	StateCSI
)

var (
	input [consoleBufLen]byte
	state = StateInput

	commands = map[string]cmdfunc{
		"":      noop,
		"ls":    ls,
		"lsblk": lsblk,
		"load":  load,
		"save":  save,
		"rm":    rm,
		"run":   run,
		"halt":  halt,
		"ping":  ping,
	}
)

func cli() {
	prompt()
	clearDisplay()
	displayMessage(2, 10, "ready")

	for i := 0; ; {
		if console.Buffered() > 0 {
			data, _ := console.ReadByte()
			switch state {
			case StateInput:
				switch data {
				case 0x8:
					fallthrough
				case 0x7f: // this is probably wrong... works on my machine tho :)
					// backspace
					if i > 0 {
						i -= 1
						console.Write([]byte{0x8, 0x20, 0x8})
					}
				case 13:
					// return key
					if console.Buffered() > 0 {
						data, _ := console.ReadByte()
						if data != 10 {
							println("\r\nunexpected: \r", int(data))
						}
					}
					console.Write([]byte("\r\n"))
					runCommand(string(input[:i]))
					prompt()

					i = 0
					continue
				case 27:
					// escape
					state = StateEscape
				default:
					// anything else, just echo the character if it is printable
					if i < (consoleBufLen - 1) {
						console.WriteByte(data)
						input[i] = data
						i++
					}
				}
			case StateEscape:
				switch data {
				case 0x5b:
					state = StateEscBrc
				default:
					state = StateInput
				}
			default:
				// TODO: handle escape sequences
				state = StateInput
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func runCommand(line string) {
	defer func() {
		p := recover()
		if p != nil {
			println("panic:", p)
		}
	}()

	argv := strings.SplitN(strings.TrimSpace(line), " ", -1)
	cmd := argv[0]
	cmdfn, ok := commands[cmd]
	if !ok {
		println("unknown command: " + line)
		return
	}
	cmdfn(argv)
}

func prompt() {
	print("==> ")
}

type cmdfunc func(argv []string)

func noop(argv []string) {}

func ls(argv []string) {
	if eng.FileStore == nil {
		println("no file store available")
		return
	}

	list, err := eng.FileStore.List()
	if err != nil {
		println("error listing files:", err.Error())
		return
	}

	print(
		"\n-------------------------------------\r\n" +
			" File Store:  \r\n" +
			"-------------------------------------\r\n")
	for _, file := range list {
		println(file.Size(), file.Name())
	}
	print(
		"\n-------------------------------------\r\n\r\n")
}

func lsblk(argv []string) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(dataStart()))
	start := hex.EncodeToString(b)

	binary.BigEndian.PutUint32(b, uint32(dataEnd()))
	end := hex.EncodeToString(b)

	print(
		"\n-------------------------------------\r\n" +
			" Device Information:  \r\n" +
			"-------------------------------------\r\n" +
			" flash data start: 0x" + start + "\r\n" +
			" flash data end:   0x" + end + "\r\n" +
			"-------------------------------------\r\n\r\n")
}

// load module into engine
func load(argv []string) {
	if eng.FileStore == nil {
		println("no file store available")
		return
	}

	if len(argv) != 2 {
		println("usage: save <name>")
		return
	}

	if running {
		println("already running. halt first.")
		return
	}

	println("loading", argv[1])
	clearDisplay()
	displayMessage(2, 10, "loading")
	displayMessage(2, 30, argv[1])

	n, err := eng.FileStore.FileSize(argv[1])
	if err != nil {
		println("error loading file:", err.Error())
		return
	}

	data := make([]byte, n)
	if err := eng.FileStore.Load(argv[1], data); err != nil {
		println("error loading file:", err.Error())
		return
	}

	if err := eng.Interpreter.Load(bytes.NewReader(data)); err != nil {
		println(err.Error())
		return
	}
	println("module loaded.")
	clearDisplay()
	displayMessage(2, 10, argv[1])
	displayMessage(2, 30, "loaded.")
}

// save into filestore
func save(argv []string) {
	if eng.FileStore == nil {
		println("no file store available")
		return
	}

	if len(argv) != 3 {
		println("usage: save <name> <size>")
		return
	}

	clearDisplay()
	displayMessage(2, 10, "saving")
	displayMessage(2, 30, argv[1])

	// read in size bytes from port
	sz := convert.StringToInt(argv[2])

	data := make([]byte, sz)
	if err := readDataFromPort(data); err != nil {
		println("error reading data:", err.Error())
		return
	}

	if err := eng.FileStore.Save(argv[1], data); err != nil {
		println("error saving file:", err.Error())
	}

	clearDisplay()
	displayMessage(2, 10, argv[1])
	displayMessage(2, 30, "saved.")
}

func readDataFromPort(data []byte) (err error) {
	for i := 0; i < len(data); i++ {
		for console.Buffered() == 0 {
		}
		data[i], err = console.ReadByte()
		if err != nil {
			return
		}
	}
	return nil
}

// remove from filestore
func rm(argv []string) {
	if eng.FileStore == nil {
		println("no file store available")
		return
	}

	if len(argv) != 2 {
		println("usage: rm <name>")
		return
	}

	clearDisplay()
	displayMessage(2, 10, "removing")
	displayMessage(2, 30, argv[1])

	if err := eng.FileStore.Remove(argv[1]); err != nil {
		println("error removing file:", err.Error())
		return
	}

	println(argv[1], "deleted.")
	clearDisplay()
	displayMessage(2, 10, argv[1])
	displayMessage(2, 30, "deleted.")
}

var (
	instance engine.Instance
	running  bool
)

func run(argv []string) {
	if eng.FileStore == nil {
		println("no file store available")
		return
	}

	if running {
		println("module already running. run 'halt' first.")
		return
	}

	// run the module
	var err error
	instance, err = eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return
	}

	running = true
	println("module running.")
	clearDisplay()
	displayMessage(2, 10, "module")
	displayMessage(2, 30, "running.")
}

func halt(argv []string) {
	if !running {
		println("module not running")
		return
	}

	println("halting...")
	_ = eng.Interpreter.Halt()
	instance = nil
	running = false
	runtime.GC()
	println("module halted.")
	clearDisplay()
	displayMessage(2, 10, "module")
	displayMessage(2, 30, "halted.")
}

func ping(argv []string) {
	if eng.FileStore == nil {
		println("no file store available")
		return
	}

	if !running {
		println("module not running. use 'run' first.")
		return
	}

	if len(argv) < 2 {
		println("usage: ping <count>")
		return
	}
	count := convert.StringToInt(argv[1])

	clearDisplay()
	displayMessage(2, 10, "ping "+argv[1])

	for i := 0; i < count; i++ {
		println("Ping...")
		if _, err := instance.Call("ping"); err != nil {
			println(err.Error())
			return
		}
	}

	displayMessage(2, 30, "done.")
}
