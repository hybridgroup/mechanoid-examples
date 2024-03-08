package main

import (
	"runtime"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"tinygo.org/x/drivers/pixel"
)

func runWASM[T pixel.Color](module string, d *display.Device[T]) error {
	println("Running WASM module", module)

	mechanoid.DebugMemory("start runWASM")
	runtime.GC()
	mechanoid.DebugMemory("start runWASM after GC")

	f, err := modules.Open("modules/" + module)
	if err != nil {
		return err
	}

	if err := eng.Interpreter.Load(f.(engine.Reader)); err != nil {
		println(err.Error())
		return err
	}

	f.Close()

	println("Running module...")
	ins, err := eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return err
	}

	if _, err := ins.Call("start"); err != nil {
		println(err.Error())
	}

	for {
		if _, err := ins.Call("update"); err != nil {
			println(err.Error())
		}

		board.Buttons.ReadInput()
		event := board.Buttons.NextEvent()
		if !event.Pressed() {
			continue
		}
		switch event.Key() {
		case board.KeySelect:
			return nil
		}

		d.Screen.Update()
		time.Sleep(time.Second / 30)
	}

	return nil
}
