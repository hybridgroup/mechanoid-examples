package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/devices/hardware"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
)

//go:embed modules/blink.wasm
var wasmModule []byte

func main() {
	time.Sleep(3 * time.Second)

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()

	intp := interp.NewInterpreter()
	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	eng.AddDevice(hardware.NewGPIODevice(eng))

	println("Initializing engine...")
	eng.Init()

	println("Loading module...")
	if err := eng.Interpreter.Load(bytes.NewReader(wasmModule)); err != nil {
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
