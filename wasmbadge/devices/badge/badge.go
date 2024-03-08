package badge

import (
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/pixel"
)

type Badge[T pixel.Color] struct {
	Display *display.Device[T]
}

func NewDevice[T pixel.Color]() *Badge[T] {
	return &Badge[T]{}
}

func (b *Badge[T]) Init() error {
	return nil
}

func (b *Badge[T]) Modules() wypes.Modules {
	return wypes.Modules{
		"badge": wypes.Module{
			"new_big_text": wypes.H1(b.newBigText),
		},
		"bigtext": wypes.Module{
			"set_text1": wypes.H2(b.bigTextSetText1),
			"set_text2": wypes.H2(b.bigTextSetText2),
			"show":      wypes.H1(b.bigTextShow),
		},
	}
}

func (b *Badge[T]) UseDisplay(d *display.Device[T]) error {
	b.Display = d
	return nil
}

func (b *Badge[T]) newBigText(msg wypes.String) wypes.HostRef[*BigText[T]] {
	// create the badge UI element
	bt := NewBigText[T](b.Display, "WASM Badge", msg.Unwrap(), "")
	if bt == nil {
		return wypes.HostRef[*BigText[T]]{Raw: nil}
	}
	bt.Show(b.Display)
	return wypes.HostRef[*BigText[T]]{Raw: bt}
}

func (b *Badge[T]) bigTextSetText1(ref wypes.HostRef[*BigText[T]], msg wypes.String) wypes.UInt32 {
	// get the badge UI element by reference
	bt := ref.Unwrap()
	if bt == nil {
		return 0
	}
	bt.SetText1(msg.Unwrap())
	return wypes.UInt32(len(msg.Unwrap()))
}

func (b *Badge[T]) bigTextSetText2(ref wypes.HostRef[*BigText[T]], msg wypes.String) wypes.UInt32 {
	// get the badge UI element by reference
	bt := ref.Unwrap()
	if bt == nil {
		return 0
	}

	bt.SetText2(msg.Unwrap())
	return wypes.UInt32(len(msg.Unwrap()))
}

func (b *Badge[T]) bigTextShow(ref wypes.HostRef[*BigText[T]]) wypes.UInt32 {
	// get the badge UI element by reference
	bt := ref.Unwrap()
	if bt == nil {
		return 0
	}
	bt.Show(b.Display)

	return 1
}
