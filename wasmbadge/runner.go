package main

import (
	"runtime"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/badge"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"tinygo.org/x/drivers/pixel"
)

func runWASM[T pixel.Color](module string, d *display.Device[T], b *badge.Badge[T]) error {
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

	// clear badge text
	b.Clear()
	b.Heading(module)

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
		case board.KeySelect, board.KeyEscape, board.KeyB:
			return nil
		}

		d.Screen.Update()
		time.Sleep(time.Second / 30)
	}

	return nil
}

func runWASMRotation[T pixel.Color](d *display.Device[T], b *badge.Badge[T]) error {
	mods, err := modules.ReadDir("modules")
	if err != nil {
		return err
	}

	i := len(mods)
	for {
		i++
		if i >= len(mods) {
			i = 0
		}

		mod := mods[i]

		if mod.IsDir() {
			continue
		}

		println("Running WASM module", mod.Name())

		mechanoid.DebugMemory("start runWASM")
		runtime.GC()
		mechanoid.DebugMemory("start runWASM after GC")

		f, err := modules.Open("modules/" + mod.Name())
		if err != nil {
			return err
		}

		if err := eng.Interpreter.Load(f.(engine.Reader)); err != nil {
			println(err.Error())
			return err
		}

		f.Close()

		println("Running module", mod.Name(), "...")
		ins, err := eng.Interpreter.Run()
		if err != nil {
			println(err.Error())
			return err
		}

		// clear badge text
		b.Clear()
		b.Heading(mod.Name())

		if _, err := ins.Call("start"); err != nil {
			println(err.Error())
		}

		start := time.Now()
		for {
			if time.Since(start) > 10*time.Second {
				break
			}

			if _, err := ins.Call("update"); err != nil {
				println(err.Error())
			}

			board.Buttons.ReadInput()
			event := board.Buttons.NextEvent()
			if !event.Pressed() {
				continue
			}
			switch event.Key() {
			case board.KeySelect, board.KeyEscape, board.KeyB:
				// return home
				return nil
			case board.KeyA:
				// skip to next
				break
			}

			d.Screen.Update()
			time.Sleep(time.Second / 10)
		}

		ins = nil
		eng.Interpreter.Halt()
	}

	return nil
}
