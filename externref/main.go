package main

import (
	"bytes"
	_ "embed"
	"time"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wasman"
	"github.com/orsinium-labs/wypes"
)

//go:embed modules/hollaback.wasm
var wasmModule []byte

var eng *engine.Engine

func main() {
	time.Sleep(3 * time.Second)

	println("Mechanoid engine starting...")
	eng = engine.NewEngine()

	intp := &wasman.Interpreter{
		Memory: make([]byte, 65536),
	}

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
			"new":   wypes.H2(newGreeter),
			"hello": wypes.H3(hello),
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
	ins.Call("start")

	for {
		println("Calling update...")
		ins.Call("update")

		time.Sleep(1 * time.Second)
	}
}

type greeter struct {
	greeting string
}

func newGreeter(ptr wypes.UInt32, sz wypes.UInt32) wypes.UInt32 {
	println("newGreeter", ptr.Unwrap(), sz.Unwrap())
	msg, err := eng.Interpreter.MemoryData(ptr.Unwrap(), sz.Unwrap())
	if err != nil {
		println(err.Error())
		return 0
	}

	println("newGreeter msg is", string(msg))

	// create the badge UI element
	g := greeter{
		greeting: string(msg),
	}

	id := wypes.UInt32(eng.Interpreter.References().Add(&g))
	println("newGreeter id is", id.Unwrap())
	return id
}

func hello(ref, ptr wypes.UInt32, sz wypes.UInt32) wypes.UInt32 {
	msg, err := eng.Interpreter.MemoryData(ptr.Unwrap(), sz.Unwrap())
	if err != nil {
		println(err.Error())
		return 0
	}

	p := eng.Interpreter.References().Get(int32(ref.Unwrap()))
	if p == nil {
		println("greet: reference not found", ref.Unwrap())
		return 0
	}

	g := p.(*greeter)
	g.greeting = string(msg)

	return sz
}
