//go:build tinygo

package main

//go:wasmimport display heading
func heading(ptr, sz uint32)

//go:wasmimport display message1
func message1(ptr, sz uint32)

//go:wasmimport display message2
func message2(ptr, sz uint32)
