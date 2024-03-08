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

	intp := interp.NewInterpreter()
	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	println("Use file store...")
	eng.UseFileStore(fs)

	println("Initializing engine...")
	err := eng.Init()
	if err != nil {
		println(err.Error())
		return
	}

	modules := wypes.Modules{
		"hosted": wypes.Module{
			"pong": wypes.H0(pongFunc),
		},
		"env": wypes.Module{
			"hola": wypes.H1(holaFunc),
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

func holaFunc(msg wypes.String) wypes.Void {
	println(msg.Unwrap())

	return wypes.Void{}
}
