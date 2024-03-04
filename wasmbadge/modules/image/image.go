package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

//go:wasmimport image new
func new_image(ptr, sz int32) int32

//go:wasmimport image show
func image_show(p int32) int32

type Image struct {
	Ref int32
}

func NewImage(data []byte) *Image {
	ptr, sz := convert.BytesToWasmPtr(data)
	return &Image{
		Ref: new_image(int32(ptr), int32(sz)),
	}
}

func (i *Image) Show() {
	image_show(i.Ref)
}
