package main

import (
	"runtime"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"tinygo.org/x/drivers/pixel"
)

func runWASM[T pixel.Color](module string, d *display.Device[T]) error {
	println("Running WASM module", module)
	ms := runtime.MemStats{}

	runtime.ReadMemStats(&ms)
	println("Heap start runWASM Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)

	f, err := modules.Open("modules/" + module)
	if err != nil {
		return err
	}

	if err := eng.Interpreter.Load(f.(engine.Reader)); err != nil {
		println(err.Error())
		return err
	}

	f.Close()

	runtime.ReadMemStats(&ms)
	println("Heap after interp load runWASM Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)

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
			ins = nil

			runtime.ReadMemStats(&ms)
			println("Heap exit runWASM Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)

			return nil
		}

		d.Screen.Update()
		time.Sleep(time.Second / 30)
	}

	return nil
}
