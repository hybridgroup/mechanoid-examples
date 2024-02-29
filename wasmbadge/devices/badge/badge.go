package badge

import (
	"unsafe"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
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

	if err := b.Engine.Interpreter.DefineFunc(moduleName, "new_big_text", b.newBigText); err != nil {
		println(err.Error())
		return err
	}

	if err := b.Engine.Interpreter.DefineFunc("bigtext", "set_text1", b.bigTextSetText1); err != nil {
		println(err.Error())
		return err
	}

	if err := b.Engine.Interpreter.DefineFunc("bigtext", "set_text2", b.bigTextSetText2); err != nil {
		println(err.Error())
		return err
	}

	if err := b.Engine.Interpreter.DefineFunc("bigtext", "show", b.bigTextShow); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func (b *Badge[T]) UseDisplay(d *display.Device[T]) error {
	b.Display = d
	return nil
}

func (b *Badge[T]) newBigText(ptr uint32, sz uint32) uint32 {
	msg, err := b.Engine.Interpreter.MemoryData(ptr, sz)
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

	id := uint32(b.Engine.Interpreter.References().Add(unsafe.Pointer(bt)))
	return id
}

func (b *Badge[T]) bigTextSetText1(ref uint32, ptr uint32, sz uint32) uint32 {
	msg, err := b.Engine.Interpreter.MemoryData(ptr, sz)
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

func (b *Badge[T]) bigTextSetText2(ref uint32, ptr uint32, sz uint32) uint32 {
	msg, err := b.Engine.Interpreter.MemoryData(ptr, sz)
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

func (b *Badge[T]) bigTextShow(ref uint32) uint32 {
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
