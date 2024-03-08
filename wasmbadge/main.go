package main

import (
	"embed"
	"runtime"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
	"tinygo.org/x/drivers/pixel"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/badge"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
)

//go:embed modules/*.wasm
var modules embed.FS

var (
	eng *engine.Engine
)

func main() {
	if board.Name == "simulator" {
		// Use the configuration for the Gopher Badge.
		board.Simulator.WindowWidth = 320
		board.Simulator.WindowHeight = 240
		board.Simulator.WindowPPI = 166
		board.Simulator.WindowDrawSpeed = time.Second * 16 / 62_500e3 // 62.5MHz, 16bpp
	}

	run(board.Display.Configure())
}

func run[T pixel.Color](disp board.Displayer[T]) {
	time.Sleep(2 * time.Second)

	mechanoid.DebugMemory("start of program")

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	intp := interp.NewInterpreter()
	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	d := display.NewDevice[T](disp)
	d.Init()

	// badge interface to display API
	bg := badge.NewDevice[T]()
	bg.UseDisplay(&d)

	// host interface to badge API
	eng.AddDevice(bg)

	println("Initializing engine...")
	eng.Init()

	mechanoid.DebugMemory("after engine init")

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
			home = nil
			runWASM(menuChoices[listbox.Selected()], &d)

			mechanoid.DebugMemory("after runWASM exit")

			bg.Cleanup()
			runtime.GC()

			mechanoid.DebugMemory("after runWASM exit GC")

			home = createHome[T](&d)
			home.Show(&d)
			listbox = home.ListBox
		}

		d.Screen.Update()
		time.Sleep(time.Second / 30)
	}
}
