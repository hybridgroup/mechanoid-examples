package badge

import (
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/pixel"
)

type Badge[T pixel.Color] struct {
	Display *display.Device[T]
	bt      *BigText[T]
}

func NewDevice[T pixel.Color](d *display.Device[T]) *Badge[T] {
	return &Badge[T]{
		Display: d,
	}
}

func (b *Badge[T]) Init() error {
	b.bt = NewBigText[T](b.Display)
	b.bt.Show(b.Display)

	return nil
}

func (b *Badge[T]) Modules() wypes.Modules {
	return wypes.Modules{
		"badge": wypes.Module{
			"heading":   wypes.H1(b.bigTextHeading),
			"set_text1": wypes.H1(b.bigTextSetText1),
			"set_text2": wypes.H1(b.bigTextSetText2),
		},
	}
}

func (b *Badge[T]) Clear() error {
	b.bt.Heading("")
	b.bt.SetText1("")
	b.bt.SetText2("")
	b.bt.Show(b.Display)

	return nil
}

func (b *Badge[T]) bigTextHeading(msg wypes.String) wypes.Void {
	b.Heading(msg.Unwrap())

	return wypes.Void{}
}

func (b *Badge[T]) bigTextSetText1(msg wypes.String) wypes.Void {
	b.SetText1(msg.Unwrap())

	return wypes.Void{}
}

func (b *Badge[T]) bigTextSetText2(msg wypes.String) wypes.Void {
	b.SetText2(msg.Unwrap())

	return wypes.Void{}
}

func (b *Badge[T]) Heading(msg string) {
	b.bt.Heading(msg)
	b.bt.Show(b.Display)
}

func (b *Badge[T]) SetText1(msg string) {
	b.bt.SetText1(msg)
	b.bt.Show(b.Display)
}

func (b *Badge[T]) SetText2(msg string) {
	b.bt.SetText2(msg)
	b.bt.Show(b.Display)
}
