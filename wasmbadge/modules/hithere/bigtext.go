package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

//go:wasmimport badge new_big_text
func new_big_text(ptr, sz uint32) uint32

//go:wasmimport bigtext set_text1
func big_text_set_text1(p uint32, ptr, sz uint32) uint32

//go:wasmimport bigtext set_text2
func big_text_set_text2(p uint32, ptr, sz uint32) uint32

//go:wasmimport bigtext show
func big_text_show(p uint32) uint32

type BigText struct {
	Ref uint32
}

func NewBigText(msg string) *BigText {
	ptr, sz := convert.StringToWasmPtr(msg)
	return &BigText{
		Ref: new_big_text(ptr, sz),
	}
}

func (b *BigText) SetText1(msg []byte) {
	ptr, sz := convert.BytesToWasmPtr(msg)
	big_text_set_text1(b.Ref, ptr, sz)
}

func (b *BigText) SetText2(msg []byte) {
	ptr, sz := convert.BytesToWasmPtr(msg)
	big_text_set_text2(b.Ref, ptr, sz)
}

func (b *BigText) Show() {
	big_text_show(b.Ref)
}
