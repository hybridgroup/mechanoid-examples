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
	// ref is an externref that is returned from the new_greeter host function.
	ref uint32

	// buf is a buffer that is used to pass messages to the greeter host instance.
	buf [64]byte
)

//go:export start
func start() {
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	// ref is an externref that is passed to the hello function.
	// this is an opaque reference to the greeter instance on the host.
	// The host is responsible for managing the lifetime of the instance.
	// It is not a pointer to the instance, but a reference to it.
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
