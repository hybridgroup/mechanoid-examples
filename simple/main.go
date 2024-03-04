package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
	"github.com/orsinium-labs/wypes"
)

//go:embed modules/ping.wasm
var pingModule []byte

func main() {
	time.Sleep(1 * time.Second)

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()

	intp := &wasman.Interpreter{}

	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	println("Initializing engine...")
	err := eng.Init()
	if err != nil {
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

	println("Loading WASM module...")
	if err := eng.Interpreter.Load(bytes.NewReader(pingModule)); err != nil {
		println(err.Error())
		return
	}

	println("Running module...")
	ins, err := eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return
	}

	for {
		println("Calling ping...")
		_, _ = ins.Call("ping")

		time.Sleep(1 * time.Second)
	}
}

func pongFunc() wypes.Void {
	println("pong")
	return wypes.Void{}
}
