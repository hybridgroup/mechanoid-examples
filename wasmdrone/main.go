package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid-examples/wasmdrone/devices/display"
	"github.com/hybridgroup/mechanoid-examples/wasmdrone/devices/drone"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
	"tinygo.org/x/drivers/pixel"
)

//go:embed modules/drone.wasm
var wasmCode []byte

var (
	eng *engine.Engine
)

var ssid, pass string

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
	dsp := display.NewDevice(disp)
	eng.AddDevice(dsp)

	tello := drone.NewDevice(ssid, pass)
	eng.AddDevice(tello)

	println("Initializing engine using interpreter", eng.Interpreter.Name())
	if err := eng.Init(); err != nil {
		println(err.Error())
		return
	}

	board.Buttons.Configure()

	dsp.Heading("WASM Drone")
	dsp.Message1Text.SetText("Loading code...")

	println("Loading and running WASM code...")
	ins, err := eng.LoadAndRun(bytes.NewReader(wasmCode))
	if err != nil {
		println(err.Error())
		return
	}

	dsp.Message1Text.SetText("Starting drone...")
	println("Starting drone...")

	if err := tello.Start(); err != nil {
		println(err.Error())
		return
	}
	go tello.Control()

	dsp.Message1Text.SetText("Ready")

	shifted := int32(0)
	for {
		board.Buttons.ReadInput()
		event := board.Buttons.NextEvent()
		// handled "shifting" using A
		if event.Key() == board.KeyA && event.Pressed() {
			shifted = 1
			continue
		} else if event.Key() == board.KeyA && !event.Pressed() {
			shifted = 0
			continue
		}

		if !event.Pressed() {
			// key released, so set direction/speed to none
			tello.Direction = drone.DirectionNone
			tello.Speed = 0

			continue
		}
		switch event.Key() {
		// case board.KeyA:
		// 	// "shifter"
		// 	if _, err := ins.Call("button_a"); err != nil {
		// 		println(err.Error())
		// 	}
		case board.KeyStart:
			// takeoff
			if _, err := ins.Call("button_start", shifted); err != nil {
				println(err.Error())
			}
		case board.KeyB:
			// lander
			if _, err := ins.Call("button_b", shifted); err != nil {
				println(err.Error())
			}
		case board.KeySelect:
			// flip
			if _, err := ins.Call("button_select", shifted); err != nil {
				println(err.Error())
			}
		case board.KeyUp:
			if _, err := ins.Call("button_up", shifted); err != nil {
				println(err.Error())
			}
		case board.KeyDown:
			if _, err := ins.Call("button_down", shifted); err != nil {
				println(err.Error())
			}
		case board.KeyLeft:
			if _, err := ins.Call("button_left", shifted); err != nil {
				println(err.Error())
			}
		case board.KeyRight:
			if _, err := ins.Call("button_right", shifted); err != nil {
				println(err.Error())
			}
		}

		time.Sleep(time.Second / 10)
	}
}
