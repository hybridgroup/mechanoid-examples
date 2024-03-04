package main

import (
	"embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wazero"
	"tinygo.org/x/drivers/pixel"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/badge"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
)

//go:embed modules/*.wasm
var modules embed.FS

var (
	eng  *engine.Engine
	intp engine.Interpreter
)

func main() {
	time.Sleep(1 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	intp = &wazero.Interpreter{}

	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	run(display.NewDevice(board.Display.Configure()))
}

func run[T pixel.Color](d display.Device[T]) {
	// badge interface to display API
	bg := badge.NewDevice[T](eng)
	bg.UseDisplay(&d)

	// host interface to badge API
	eng.AddDevice(bg)

	println("Initializing engine...")
	eng.Init()

	board.Buttons.Configure()

	loadMenuChoices()
	home := createHome[T](&d)
	home.Show(&d)

	listbox := home.ListBox

	for {
		board.Buttons.ReadInput()
		event := board.Buttons.NextEvent()
		if !event.Pressed() {
			continue
		}
		switch event.Key() {
		case board.KeyUp:
			index := listbox.Selected() - 1
			if index < 0 {
				index = listbox.Len() - 1
			}
			listbox.Select(index)
		case board.KeyDown:
			index := listbox.Selected() + 1
			if index >= listbox.Len() {
				index = 0
			}
			listbox.Select(index)
		case board.KeyEnter, board.KeyA:
			module := menuChoices[listbox.Selected()]
			runWASM(module, &d)
			eng.Interpreter.Halt()

			home.Show(&d)
		}

		d.Screen.Update()
		time.Sleep(time.Second / 30)
	}
}
