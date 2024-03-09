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
var wasmCode []byte

func main() {
	time.Sleep(3 * time.Second)

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()
	eng.UseInterpreter(interp.NewInterpreter())
	eng.AddDevice(hardware.GPIO{})

	println("Initializing engine using interpreter", eng.Interpreter.Name())
	if err := eng.Init(); err != nil {
		println(err.Error())
		return
	}

	println("Loading and running WASM code...")
	ins, err := eng.LoadAndRun(bytes.NewReader(wasmCode))
	if err != nil {
		println(err.Error())
		return
	}

	println("Calling setup...")
	if _, err := ins.Call("setup"); err != nil {
		println(err.Error())
	}

	for {
		println("Calling loop...")
		if _, err := ins.Call("loop"); err != nil {
			println(err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}
