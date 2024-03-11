//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

const msg = "You are at"
const name = "Wasm I/O"

//go:export start
func start() {
	badge_set_text2(convert.StringToWasmPtr(msg))
}

//go:export update
func update() {
	badge_set_text4(convert.StringToWasmPtr(name))
}

func main() {}
