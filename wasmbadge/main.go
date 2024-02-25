package main

import (
	"embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid/convert"
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

	board.Buttons.Configure()

	run(NewDisplayDevice(board.Display.Configure()))
}

func run[T pixel.Color](display DisplayDevice[T]) {
	loadMenuChoices()
	display.createHome()
	display.showHome()

	listbox := homePage.(*Page[T]).ListBox

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
			runWASM(menuChoices[listbox.Selected()], display, homePage)
			display.showHome()
		}

		display.Screen.Update()
		time.Sleep(time.Second / 30)
	}
}

func runWASM[T pixel.Color](module string, display DisplayDevice[T], home any) error {
	println("Running WASM module", module)

	display.createWasmPage(module)
	display.showWasmPage()

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

	for i := 0; i < 15; i++ {
		ins.Call("ping")
		pingCount++

		println("Ping", pingCount)
		display.outputToWasmPage("Ping " + convert.IntToString(pingCount))

		time.Sleep(1 * time.Second)
	}

	return nil
}
