package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

//go:wasmimport image new
func new_image(ptr, sz uint32) uint32

//go:wasmimport image show
func image_show(p uint32) uint32

type Image struct {
	Ref uint32
}

func NewImage(data string) *Image {
	ptr, sz := convert.StringToWasmPtr(data)
	return &Image{
		Ref: new_image(ptr, sz),
	}
}

func (i *Image) Show() {
	image_show(i.Ref)
}
