//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

var (
	// buf is a buffer that is used to pass messages to the greeter host instance.
	buf [64]byte
)

//go:wasmimport env hola
func hola(ptr, sz uint32)

const msg = "Hello, WebAssembly!"

//go:export hello
func hello() {
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	hola(ptr, sz)
}

func main() {}
