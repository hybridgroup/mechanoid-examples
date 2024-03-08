package main

import (
	"bytes"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"tinygo.org/x/drivers/pixel"
)

func runWASM[T pixel.Color](module string, d *display.Device[T]) error {
	println("Running WASM module", module)

	moduleData, err := modules.ReadFile("modules/" + module)
	if err != nil {
		return err
	}

	if err := eng.Interpreter.Load(bytes.NewReader(moduleData)); err != nil {
		println(err.Error())
		return err
	}

	println("Running module...")
	ins, err := eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return err
	}

	ins.Call("start")

	for {
		ins.Call("update")

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
