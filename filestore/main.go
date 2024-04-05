package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
	"github.com/orsinium-labs/wypes"
)

var (
	eng *engine.Engine
)

func main() {
	time.Sleep(1 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()
	eng.UseInterpreter(interp.NewInterpreter())
	eng.UseFileStore(fs)

	println("Initializing engine using interpreter", eng.Interpreter.Name())
	if err := eng.Init(); err != nil {
		println(err.Error())
		return
	}

	modules := wypes.Modules{
		"hosted": wypes.Module{
			"pong": wypes.H0(pongFunc),
		},
	}
	if err := eng.Interpreter.SetModules(modules); err != nil {
		println(err.Error())
		return
	}
	// start up CLI
	cli()
}

func pongFunc() wypes.Void {
	println("pong")
	return wypes.Void{}
}
