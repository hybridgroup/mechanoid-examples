package badge

import (
	"unsafe"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/pixel"
)

type Badge[T pixel.Color] struct {
	Engine  *engine.Engine
	Display *display.Device[T]
}

func NewDevice[T pixel.Color](e *engine.Engine) *Badge[T] {
	return &Badge[T]{
		Engine: e,
	}
}

func (b *Badge[T]) Init() error {
	if b.Engine == nil {
		return engine.ErrInvalidEngine
	}

	modules := wypes.Modules{
		"badge": wypes.Module{
			"new_big_text": wypes.H2(b.newBigText),
		},
		"bigtext": wypes.Module{
			"set_text1": wypes.H3(b.bigTextSetText1),
			"set_text2": wypes.H3(b.bigTextSetText2),
			"show":      wypes.H1(b.bigTextShow),
		},
	}
	if err := b.Engine.Interpreter.SetModules(modules); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func (b *Badge[T]) UseDisplay(d *display.Device[T]) error {
	b.Display = d
	return nil
}

func (b *Badge[T]) newBigText(ptr wypes.UInt32, sz wypes.UInt32) wypes.UInt32 {
	msg, err := b.Engine.Interpreter.MemoryData(ptr.Unwrap(), sz.Unwrap())
	if err != nil {
		println(err.Error())
		return 0
	}

	// create the badge UI element
	bt := NewBigText[T](b.Display, "WASM Badge", string(msg), "")
	if bt == nil {
		return 0
	}
	bt.Show(b.Display)

	id := wypes.UInt32(b.Engine.Interpreter.References().Add(unsafe.Pointer(bt)))
	return id
}

func (b *Badge[T]) bigTextSetText1(ref wypes.UInt32, ptr wypes.UInt32, sz wypes.UInt32) wypes.UInt32 {
	msg, err := b.Engine.Interpreter.MemoryData(ptr.Unwrap(), sz.Unwrap())
	if err != nil {
		println(err.Error())
		return 0
	}

	// get the badge UI element by reference
	p := b.Engine.Interpreter.References().Get(int32(ref))
	if p == uintptr(0) {
		println("bigTextSetText1: reference not found")
		return 0
	}
	bt := (*BigText[T])(unsafe.Pointer(p))
	if bt == nil {
		return 0
	}
	bt.SetText1(string(msg))

	return sz
}

func (b *Badge[T]) bigTextSetText2(ref wypes.UInt32, ptr wypes.UInt32, sz wypes.UInt32) wypes.UInt32 {
	msg, err := b.Engine.Interpreter.MemoryData(ptr.Unwrap(), sz.Unwrap())
	if err != nil {
		println(err.Error())
		return 0
	}

	// get the badge UI element by reference
	p := b.Engine.Interpreter.References().Get(int32(ref))
	if p == uintptr(0) {
		println("bigTextSetText2: reference not found")
		return 0
	}
	bt := (*BigText[T])(unsafe.Pointer(p))
	if bt == nil {
		return 0
	}

	bt.SetText2(string(msg))

	return sz
}

func (b *Badge[T]) bigTextShow(ref wypes.UInt32) wypes.UInt32 {
	// get the badge UI element by reference
	p := b.Engine.Interpreter.References().Get(int32(ref))
	if p == uintptr(0) {
		println("bigTextSetText1: reference not found")
		return 0
	}
	bt := (*BigText[T])(unsafe.Pointer(p))
	if bt == nil {
		return 0
	}
	bt.Show(b.Display)

	return 1
}
