package main

import (
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
	"github.com/orsinium-labs/wypes"
)

var (
	eng *engine.Engine
)

func main() {
	time.Sleep(1 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	println("Using interpreter...")
	eng.UseInterpreter(&wasman.Interpreter{
		Memory: make([]byte, 65536),
	})

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
			"hola": wypes.H2(holaFunc),
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

func holaFunc(ptr wypes.UInt32, size wypes.UInt32) wypes.UInt32 {
	msg, err := eng.Interpreter.MemoryData(ptr.Unwrap(), size.Unwrap())
	if err != nil {
		println(err.Error())
		return 0
	}
	println(string(msg))
	return size
}
