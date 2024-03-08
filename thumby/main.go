package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid-examples/thumby/devices/display"
	"github.com/hybridgroup/mechanoid/convert"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
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

	intp := interp.NewInterpreter()
	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	disp := &display.Device{}
	eng.AddDevice(disp)

	println("Initializing engine...")
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
		pingcount++
		println("Calling ping", pingcount)

		disp.Clear()
		disp.ShowMessage(5, 10, "ping "+convert.IntToString(pingcount))

		if _, err := ins.Call("ping"); err != nil {
			println(err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}

func pongFunc() wypes.Void {
	pongcount++

	println("pong", pongcount)

	eng.Devices[0].(*display.Device).ShowMessage(5, 30, "pong "+convert.IntToString(pongcount))

	return wypes.Void{}
}
