package badge

import (
	"context"
	"unsafe"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"github.com/hybridgroup/mechanoid/interp/wazero"
	"github.com/stealthrocket/wazergo"
	"github.com/stealthrocket/wazergo/types"
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
	interp := b.Engine.Interpreter.(*wazero.Interpreter)
	err := wazero.AddModule[*Badge[T]](interp, "badge", b, wazergo.Functions[*Badge[T]]{
		"new_big_text": wazergo.F1((*Badge[T]).newBigText),
	})
	if err != nil {
		println(err.Error())
		return err
	}
	err = wazero.AddModule[*Badge[T]](interp, "bigtext", b, wazergo.Functions[*Badge[T]]{
		"set_text1": wazergo.F2((*Badge[T]).bigTextSetText1),
		"set_text2": wazergo.F2((*Badge[T]).bigTextSetText2),
		"show":      wazergo.F1((*Badge[T]).bigTextShow),
	})
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func (b *Badge[T]) UseDisplay(d *display.Device[T]) error {
	b.Display = d
	return nil
}

func (b *Badge[T]) newBigText(_ context.Context, msg types.String) types.Uint32 {
	// create the badge UI element
	bt := NewBigText[T](b.Display, "WASM Badge", string(msg), "")
	if bt == nil {
		return 0
	}
	bt.Show(b.Display)

	id := uint32(b.Engine.Interpreter.References().Add(unsafe.Pointer(bt)))
	return types.Uint32(id)
}

func (b *Badge[T]) bigTextSetText1(_ context.Context, ref types.Uint32, msg types.String) types.Uint32 {
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
	return types.Uint32(len(msg))
}

func (b *Badge[T]) bigTextSetText2(_ context.Context, ref types.Uint32, msg types.String) types.Uint32 {
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
	return types.Uint32(len(msg))
}

func (b *Badge[T]) bigTextShow(_ context.Context, ref types.Uint32) types.Uint32 {
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

func (b *Badge[T]) Close(context.Context) error {
	return nil
}
