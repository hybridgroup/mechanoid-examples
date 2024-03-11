//go:build tinygo

package main

//go:wasmimport badge heading
func badge_heading(ptr, sz uint32)

//go:wasmimport badge set_text1
func badge_set_text1(ptr, sz uint32)

//go:wasmimport badge set_text2
func badge_set_text2(ptr, sz uint32)
