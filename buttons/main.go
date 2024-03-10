package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid-examples/buttons/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
	"tinygo.org/x/drivers/pixel"
)

//go:embed modules/buttons.wasm
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

	board.Buttons.Configure()

	println("Loading and running WASM code...")
	ins, err := eng.LoadAndRun(bytes.NewReader(wasmCode))
	if err != nil {
		println(err.Error())
		return
	}

	for {
		board.Buttons.ReadInput()
		event := board.Buttons.NextEvent()
		if !event.Pressed() {
			continue
		}
		switch event.Key() {
		case board.KeyA:
			if _, err := ins.Call("button_a"); err != nil {
				println(err.Error())
			}
		case board.KeyB:
			if _, err := ins.Call("button_b"); err != nil {
				println(err.Error())
			}
		case board.KeyUp:
			if _, err := ins.Call("button_up"); err != nil {
				println(err.Error())
			}
		case board.KeyDown:
			if _, err := ins.Call("button_down"); err != nil {
				println(err.Error())
			}
		case board.KeyLeft:
			if _, err := ins.Call("button_left"); err != nil {
				println(err.Error())
			}
		case board.KeyRight:
			if _, err := ins.Call("button_right"); err != nil {
				println(err.Error())
			}
		case board.KeySelect, board.KeyEscape:
			if _, err := ins.Call("button_select"); err != nil {
				println(err.Error())
			}
		case board.KeyStart, board.KeyEnter:
			if _, err := ins.Call("button_start"); err != nil {
				println(err.Error())
			}
		}

		time.Sleep(time.Second / 30)
	}
}
