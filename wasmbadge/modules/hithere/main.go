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

	badge_set_text1(convert.BytesToWasmPtr(buf[:len(msg)]))
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

	badge_set_text3(convert.BytesToWasmPtr(buf[0:end]))
}

func main() {}
