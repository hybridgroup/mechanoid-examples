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
var wasmCode []byte

var (
	eng *engine.Engine

	pingcount, pongcount int
)

func main() {
	time.Sleep(3 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()
	eng.UseInterpreter(interp.NewInterpreter())

	disp := &display.Device{}
	eng.AddDevice(disp)

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
