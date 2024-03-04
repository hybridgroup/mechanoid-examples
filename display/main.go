package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/aykevl/board"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
	"github.com/orsinium-labs/wypes"
)

//go:embed modules/ping.wasm
var pingModule []byte

var (
	pingCount, pongCount int
)

func main() {
	time.Sleep(1 * time.Second)

	display := NewDisplayDevice(board.Display.Configure())

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{})

	println("Initializing engine...")
	eng.Init()

	modules := wypes.Modules{
		"hosted": wypes.Module{
			"pong": wypes.H0(func() wypes.Void {
				pongCount++
				println("pong", pongCount)
				display.Pong(pongCount)
				return wypes.Void{}
			}),
		},
	}
	if err := eng.Interpreter.SetModules(modules); err != nil {
		println(err.Error())
		return
	}

	println("Loading module...")
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
		println("Ping", pingCount)
		ins.Call("ping")
		pingCount++
		display.Ping(pingCount)

		time.Sleep(1 * time.Second)
	}
}
