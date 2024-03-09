package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
	"github.com/orsinium-labs/wypes"
)

//go:embed modules/ping.wasm
var wasmCode []byte

func main() {
	time.Sleep(2 * time.Second)

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()
	eng.UseInterpreter(interp.NewInterpreter())

	println("Initializing engine using interpreter", eng.Interpreter.Name())
	if err := eng.Init(); err != nil {
		println(err.Error())
		return
	}

	println("Defining host function...")
	modules := wypes.Modules{
		"hosted": wypes.Module{
			"pong": wypes.H0(pongFunc),
		},
	}
	if err := eng.Interpreter.SetModules(modules); err != nil {
		println(err.Error())
		return
	}

	println("Loading and running WASM code...")
	ins, err := eng.LoadAndRun(bytes.NewReader(wasmCode))
	if err != nil {
		println(err.Error())
		return
	}

	for {
		println("Calling ping...")
		if _, err := ins.Call("ping"); err != nil {
			println(err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}

func pongFunc() wypes.Void {
	println("pong")
	return wypes.Void{}
}
