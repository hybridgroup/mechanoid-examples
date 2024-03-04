package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

//go:wasmimport bigtext new
func new_big_text(ptr, sz int32) int32

//go:wasmimport bigtext text1
func big_text_set_text1(p int32, ptr, sz int32) int32

//go:wasmimport bigtext text2
func big_text_set_text2(p int32, ptr, sz int32) int32

//go:wasmimport bigtext show
func big_text_show(p int32) int32

type BigText struct {
	Ref int32
}

func NewBigText(msg string) *BigText {
	ptr, sz := convert.StringToWasmPtr(msg)
	return &BigText{
		Ref: new_big_text(int32(ptr), int32(sz)),
	}
}

func (b *BigText) SetText1(msg []byte) {
	ptr, sz := convert.BytesToWasmPtr(msg)
	big_text_set_text1(b.Ref, int32(ptr), int32(sz))
}

func (b *BigText) SetText2(msg []byte) {
	ptr, sz := convert.BytesToWasmPtr(msg)
	big_text_set_text2(b.Ref, int32(ptr), int32(sz))
}

func (b *BigText) Show() {
	big_text_show(b.Ref)
}
