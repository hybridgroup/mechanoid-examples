package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid-examples/thumby/devices/display"
	"github.com/hybridgroup/mechanoid/convert"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wazero"
	"github.com/orsinium-labs/wypes"
)

//go:embed modules/ping.wasm
var wasmModule []byte

var (
	eng *engine.Engine

	pingcount, pongcount int
)

func main() {
	time.Sleep(3 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	intp := &wazero.Interpreter{}

	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	println("Using display")
	eng.AddDevice(display.NewDevice(eng))

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

	for {
		eng.Devices[0].(*display.Device).Clear()
		pingcount++
		println("Calling ping", pingcount)
		eng.Devices[0].(*display.Device).ShowMessage(5, 10, "ping "+convert.IntToString(pingcount))
		_, _ = ins.Call("ping")

		time.Sleep(1 * time.Second)
	}
}

func pongFunc() wypes.Void {
	pongcount++
	println("pong", pongcount)
	eng.Devices[0].(*display.Device).ShowMessage(5, 30, "pong "+convert.IntToString(pongcount))
	return wypes.Void{}
}
