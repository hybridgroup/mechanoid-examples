package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/devices/hardware"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
)

//go:embed modules/blink.wasm
var pingModule []byte

func main() {
	time.Sleep(5 * time.Second)

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{
		Memory: make([]byte, 65536),
	})

	eng.AddDevice(hardware.NewGPIODevice(eng))

	println("Initializing engine...")
	eng.Init()

	println("Loading module...")
	if err := eng.Interpreter.Load(pingModule); err != nil {
		println(err.Error())
		return
	}

	println("Running module...")
	ins, err := eng.Interpreter.Run()
	if err != nil {
		println(err.Error())
		return
	}

	println("Calling setup...")
	ins.Call("setup")

	for {
		println("Calling loop...")
		ins.Call("loop")

		time.Sleep(1 * time.Second)
	}
}
