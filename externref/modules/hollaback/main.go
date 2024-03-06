//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

//go:wasmimport greeter new
func new_greeter(ptr, sz uint32) uint32

//go:wasmimport greeter hello
func greeter_hello(ref, ptr, sz uint32)

//go:wasmimport greeter print_u32
func print_u32(x uint32)

const (
	msg  = "Hello, WebAssembly!"
	msg2 = "From Mechanoid"
)

var (
	ref uint32
	buf [64]byte
)

//go:export start
func start() {
	start, end := 0, len(msg)
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[start:end])
	ref = new_greeter(ptr, sz)
	print_u32(ref)
}

//go:export update
func update() {
	print_u32(ref)
	copy(buf[:], msg2)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg2)])
	greeter_hello(ref, ptr, sz)
}

func main() {}