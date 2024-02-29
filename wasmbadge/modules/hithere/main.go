//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

//go:wasmimport badge new_big_text
func new_big_text(ptr, sz uint32) uint32

//go:wasmimport bigtext set_text1
func big_text_set_text1(p uint32, ptr, sz uint32) uint32

//go:wasmimport bigtext set_text2
func big_text_set_text2(p uint32, ptr, sz uint32) uint32

//go:wasmimport bigtext show
func big_text_show(p uint32) uint32

const msg = "Hello, WebAssembly!"

var (
	bgtext  uint32
	counter int
	buf     [64]byte
)

//go:export start
func start() {
	ptr, sz := convert.StringToWasmPtr(msg)
	bgtext = new_big_text(ptr, sz)
}

//go:export update
func update() {
	p1, sz1 := convert.StringToWasmPtr("Hello, WebAssembly!")

	m := "Count: "
	copy(buf[:], m)
	counter++
	s := convert.IntToString(counter)
	copy(buf[len(m):], s)
	p2, sz2 := convert.BytesToWasmPtr(buf[:len(m)+len(s)])

	big_text_set_text1(bgtext, p1, sz1)
	big_text_set_text2(bgtext, p2, sz2)
	big_text_show(bgtext)
}

func main() {}
