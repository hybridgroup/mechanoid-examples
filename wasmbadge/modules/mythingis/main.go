//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

const message = "Mechanoid"
const message2 = "WASM framework"
const message3 = "for embedded devices"
const message4 = "mechanoid.io"

//go:export start
func start() {
	badge_set_text1(convert.StringToWasmPtr(message))
	badge_set_text2(convert.StringToWasmPtr(message2))
}

//go:export update
func update() {
	badge_set_text3(convert.StringToWasmPtr(message3))
	badge_set_text4(convert.StringToWasmPtr(message4))
}

func main() {}
