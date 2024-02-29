package main

import (
	"embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
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
	time.Sleep(5 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	intp = &wasman.Interpreter{
		Memory: make([]byte, 65536),
	}

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
		// TODO: wait for input instead of polling
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
			runWASM(menuChoices[listbox.Selected()], &d)
			home.Show(&d)
		}

		d.Screen.Update()
		time.Sleep(time.Second / 30)
	}
}

func runWASM[T pixel.Color](module string, d *display.Device[T]) error {
	println("Running WASM module", module)

	moduleData, err := modules.ReadFile("modules/" + module)
	if err != nil {
		return err
	}

	if err := eng.Interpreter.Load(moduleData); err != nil {
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

	for i := 0; i < 15; i++ {
		ins.Call("update")
		time.Sleep(1 * time.Second)
	}

	return nil
}
