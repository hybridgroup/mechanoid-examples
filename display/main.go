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
var wasmCode []byte

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
	time.Sleep(2 * time.Second)

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()
	eng.UseInterpreter(interp.NewInterpreter())
	eng.AddDevice(display.NewDevice(disp))

	println("Initializing engine using interpreter", eng.Interpreter.Name())
	if err := eng.Init(); err != nil {
		println(err.Error())
		return
	}

	println("Loading and running WASM code...")
	ins, err := eng.LoadAndRun(bytes.NewReader(wasmCode))
	if err != nil {
		println(err.Error())
		return
	}

	for {
		pingCount++
		println("Ping", pingCount)

		if _, err := ins.Call("ping"); err != nil {
			println(err.Error())
		}

		eng.Devices[0].(*display.Device[T]).ShowPing(pingCount)

		time.Sleep(1 * time.Second)
	}
}
