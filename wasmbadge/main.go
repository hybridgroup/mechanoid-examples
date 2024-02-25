package main

import (
	"embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
	"tinygo.org/x/drivers/pixel"
)

//go:embed modules/*.wasm
var modules embed.FS

var (
	pingCount, pongCount int

	eng  *engine.Engine
	intp engine.Interpreter
)

func main() {
	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	intp = &wasman.Interpreter{
		Memory: make([]byte, 65536),
	}

	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	println("Initializing engine...")
	eng.Init()

	if err := eng.Interpreter.DefineFunc("hosted", "pong", func() {
		pongCount++
		//display.SetText2(convert.IntToString(pongCount))
	}); err != nil {
		println(err.Error())
		return
	}

	run(NewDisplayDevice(board.Display.Configure()))
}

func loadMenuChoices() error {
	modules, err := modules.ReadDir("modules")
	if err != nil {
		return err
	}

	for _, module := range modules {
		if module.IsDir() {
			continue
		}
		menuChoices = append(menuChoices, module.Name())
	}
	return nil
}

func run[T pixel.Color](display DisplayDevice[T]) {
	loadMenuChoices()
	display.createHome()
	display.showHome()

	listbox := homePage.(*Page[T]).ListBox

	for {
		// TODO: wait for input instead of polling
		board.Buttons.ReadInput()
		for {
			event := board.Buttons.NextEvent()
			if event == board.NoKeyEvent {
				break
			}
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
				//runApp(listbox.Selected(), display, screen, home, touchInput)
			}
		}

		display.Screen.Update()
		time.Sleep(time.Second / 30)
	}
}

func runWASM(module string) error {
	println("Loading WASM module...")
	moduleData, err := modules.ReadFile("modules/" + module)

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

	for {
		println("Ping", pingCount)
		ins.Call("ping")
		pingCount++
		//display.SetText1(convert.IntToString(pingCount))

		time.Sleep(1 * time.Second)
	}
}
