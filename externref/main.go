package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp"
	"github.com/orsinium-labs/wypes"
)

//go:embed modules/hollaback.wasm
var wasmModule []byte

func main() {
	time.Sleep(3 * time.Second)

	println("Mechanoid engine starting...")
	eng := engine.NewEngine()

	intp := interp.NewInterpreter()
	println("Using interpreter", intp.Name())
	eng.UseInterpreter(intp)

	println("Initializing engine...")
	err := eng.Init()
	if err != nil {
		println(err.Error())
		return
	}

	println("Defining host function...")
	modules := wypes.Modules{
		"greeter": wypes.Module{
			"new":       wypes.H1(newGreeter),
			"hello":     wypes.H2(hello),
			"print_u32": wypes.H1(printU32),
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

	println("Calling start...")
	if _, err := ins.Call("start"); err != nil {
		println(err.Error())
	}
	for {
		println("Calling update...")
		if _, err := ins.Call("update"); err != nil {
			println(err.Error())
		}

		time.Sleep(1 * time.Second)
	}
}

type greeter struct {
	greeting string
}

func newGreeter(msg wypes.String) wypes.HostRef[greeter] {
	println("newGreeter msg is", msg.Unwrap())
	// create the badge UI element
	g := greeter{
		greeting: msg.Unwrap(),
	}
	return wypes.HostRef[greeter]{Raw: g}
}

func hello(ref wypes.HostRef[greeter], msg wypes.String) wypes.Void {
	println("hello msg is", msg.Unwrap())
	g := ref.Unwrap()
	g.greeting = msg.Unwrap()
	return wypes.Void{}
}

func printU32(x wypes.UInt32) wypes.Void {
	println("got value:", x.Unwrap())
	return wypes.Void{}
}
