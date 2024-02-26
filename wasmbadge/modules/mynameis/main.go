//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

//go:wasmimport badge new_big_text
func new_big_text(ptr, sz uint32) uint32

const msg = "TinyGo"

//go:export start
func start() {
	ptr, sz := convert.StringToWasmPtr(msg)
	new_big_text(ptr, sz)
}

//go:export update
func update() {
	// TODO: something with the ui
}

func main() {}
