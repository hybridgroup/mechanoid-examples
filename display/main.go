package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid-examples/display/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
	"tinygo.org/x/drivers/pixel"
)

//go:embed modules/ping.wasm
var wasmModule []byte

var (
	eng *engine.Engine

	pingCount, pongCount int
)

// main func just calls a run() func so we can infer the display type.
func main() {
	run(board.Display.Configure())
}

// run func is the main entry point for the program.
func run[T pixel.Color](disp board.Displayer[T]) {
	time.Sleep(1 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	println("Adding display device...")
	eng.AddDevice(display.NewDevice(eng, disp))

	intp := interp.NewInterpreter()
	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	println("Initializing engine...")
	eng.Init()

	println("Loading WASM module...")
	if err := eng.Interpreter.Load(bytes.NewReader(wasmModule)); err != nil {
		println(err.Error())
		return
	}

	println("Running module...")
	ins, err := eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return
	}

	for {
		pingCount++
		println("Ping", pingCount)
		ins.Call("ping")
		eng.Devices[0].(*display.Device[T]).ShowPing(pingCount)

		time.Sleep(1 * time.Second)
	}
}
