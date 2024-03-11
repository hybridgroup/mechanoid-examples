//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

const msg = "This badge runs WebAssembly!"
const msg2 = "Mechanoid + TinyGo"

//go:export start
func start() {
	ptr, sz := convert.StringToWasmPtr(msg)
	badge_set_text1(ptr, sz)
}

//go:export update
func update() {
	ptr, sz := convert.StringToWasmPtr(msg2)
	badge_set_text2(ptr, sz)
}

func main() {}
