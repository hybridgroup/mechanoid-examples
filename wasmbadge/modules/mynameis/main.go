//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

const name = "@mechanoidio"
const description = "WebAssembly"
const description2 = "on TinyGo"
const nada = ""

var (
	// buf is a buffer that is used to pass messages to the greeter host instance.
	buf [64]byte
)

//go:export start
func start() {
	badge_set_text1(convert.StringToWasmPtr(name))
}

//go:export update
func update() {
	copy(buf[:], description)
	badge_set_text3(convert.BytesToWasmPtr(buf[:len(description)]))

	copy(buf[len(description)+1:], description2)
	badge_set_text4(convert.BytesToWasmPtr(buf[len(description)+1 : len(description)+len(description2)+1]))
}

func main() {}
