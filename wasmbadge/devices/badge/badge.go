package badge

import (
	"github.com/aykevl/tinygl/image"
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/pixel"
)

type Badge[T pixel.Color] struct {
	Engine  *engine.Engine
	Display *display.Device[T]
}

const moduleName = "badge"

func NewDevice[T pixel.Color](e *engine.Engine) *Badge[T] {
	return &Badge[T]{
		Engine: e,
	}
}

func (b *Badge[T]) Init() error {
	if b.Engine == nil {
		return engine.ErrInvalidEngine
	}

	return b.Engine.Interpreter.SetModules(b.Modules())
}

func (b *Badge[T]) Modules() wypes.Modules {
	return wypes.Modules{
		"bigtext": {
			"new":   wypes.H2(b.newBigText),
			"text1": wypes.H3(b.bigTextSetText1),
			"text2": wypes.H3(b.bigTextSetText2),
			"show":  wypes.H1(b.bigTextShow),
		},
		"image": {
			"new":  wypes.H2(b.newImage),
			"show": wypes.H1(b.imageShow),
		},
	}
}

func (b *Badge[T]) UseDisplay(d *display.Device[T]) error {
	b.Display = d
	return nil
}

func (b *Badge[T]) newBigText(ptr wypes.Int32, sz wypes.Int32) wypes.Int32 {
	println("newBigText", ptr.Unwrap(), sz.Unwrap())
	msg, err := b.Engine.Interpreter.MemoryData(uint32(ptr.Unwrap()), uint32(sz.Unwrap()))
	if err != nil {
		println(err.Error())
		return 0
	}

	println("creating big text", string(msg))
	// create the badge UI element
	bt := NewBigText[T](b.Display, "WASM Badge", string(msg), "")
	if bt == nil {
		return 0
	}

	println("showing big text")
	bt.Show(b.Display)

	println("adding big text to references")
	return wypes.Int32(b.Engine.Interpreter.References().Add(bt))
}

func (b *Badge[T]) bigTextSetText1(ref wypes.Int32, ptr wypes.Int32, sz wypes.Int32) wypes.Int32 {
	msg, err := b.Engine.Interpreter.MemoryData(uint32(ptr.Unwrap()), uint32(sz.Unwrap()))
	if err != nil {
		println(err.Error())
		return 0
	}

	// get the badge UI element by reference
	p := b.Engine.Interpreter.References().Get(ref.Unwrap())
	if p == uintptr(0) {
		println("bigTextSetText1: reference not found")
		return 0
	}
	bt := p.(*BigText[T])
	if bt == nil {
		return 0
	}
	bt.SetText1(string(msg))

	return sz
}

func (b *Badge[T]) bigTextSetText2(ref wypes.Int32, ptr wypes.Int32, sz wypes.Int32) wypes.Int32 {
	msg, err := b.Engine.Interpreter.MemoryData(uint32(ptr.Unwrap()), uint32(sz.Unwrap()))
	if err != nil {
		println(err.Error())
		return 0
	}

	// get the badge UI element by reference
	p := b.Engine.Interpreter.References().Get(ref.Unwrap())
	if p == uintptr(0) {
		println("bigTextSetText2: reference not found")
		return 0
	}
	bt := p.(*BigText[T])
	if bt == nil {
		return 0
	}

	bt.SetText2(string(msg))

	return sz
}

func (b *Badge[T]) bigTextShow(ref wypes.Int32) wypes.Int32 {
	println("bigTextShow", ref.Unwrap())
	// get the badge UI element by reference
	p := b.Engine.Interpreter.References().Get(ref.Unwrap())
	if p == uintptr(0) {
		println("bigTextSetText1: reference not found")
		return 0
	}
	bt := p.(*BigText[T])
	if bt == nil {
		return 0
	}
	bt.Show(b.Display)

	return 1
}

func (b *Badge[T]) newImage(ptr wypes.Int32, sz wypes.Int32) wypes.Int32 {
	println("newImage", ptr, sz)
	data, err := b.Engine.Interpreter.MemoryData(uint32(ptr.Unwrap()), uint32(sz.Unwrap()))
	if err != nil {
		println(err.Error())
		return 0
	}

	// load the image
	qoi, err := image.NewQOIFromBytes[T](data)
	if err != nil {
		println(err.Error())
		return 0
	}

	// create the badge Image UI element
	img := NewImage[T](b.Display, "Image", qoi)
	if img == nil {
		return 0
	}

	return wypes.Int32(b.Engine.Interpreter.References().Add(img))
}

func (b *Badge[T]) imageShow(ref wypes.Int32) wypes.Int32 {
	// get the badge UI element by reference
	p := b.Engine.Interpreter.References().Get(ref.Unwrap())
	if p == uintptr(0) {
		println("imageShow: reference not found")
		return 0
	}
	img := p.(*Image[T]) //(unsafe.Pointer(p))
	if img == nil {
		return 0
	}
	img.Show(b.Display)

	return 1
}
