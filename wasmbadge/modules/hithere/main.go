//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

const msg = "Hello, WebAssembly!"

var (
	counter int
	buf     [64]byte
)

//go:export start
func start() {
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	badge_set_text1(ptr, sz)
}

//go:export update
func update() {
	m := "Count: "
	start, end := 0, len(m)
	copy(buf[start:end], m)

	counter++
	s := convert.IntToString(counter)
	start = end
	end = end + len(s)

	copy(buf[start:end], s)
	ptr, sz := convert.BytesToWasmPtr(buf[0:end])
	badge_set_text2(ptr, sz)
}

func main() {}
