//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

var (
	// buf is a buffer that is used to pass messages to the host instance.
	buf [64]byte
)

//go:wasmimport display message
func message(ptr, sz uint32)

//go:export button_select
func buttonSelect() {
	msg := "SELECT"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

//go:export button_start
func buttonStart() {
	msg := "START"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

//go:export button_a
func buttonA() {
	msg := "A"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

//go:export button_b
func buttonB() {
	msg := "B"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

//go:export button_up
func buttonUp() {
	msg := "UP"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

//go:export button_down
func buttonDown() {
	msg := "DOWN"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

//go:export button_left
func buttonLeft() {
	msg := "LEFT"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

//go:export button_right
func buttonRight() {
	msg := "RIGHT"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message(ptr, sz)
}

func main() {}
