//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

const msg = "This badge runs"
const msg2 = "WebAssembly!"
const msg3 = "Mechanoid + TinyGo"

//go:export start
func start() {
	badge_set_text1(convert.StringToWasmPtr(msg))
}

//go:export update
func update() {
	badge_set_text2(convert.StringToWasmPtr(msg2))

	badge_set_text4(convert.StringToWasmPtr(msg3))
}

func main() {}
