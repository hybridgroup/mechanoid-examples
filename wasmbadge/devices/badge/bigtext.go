package badge

import (
	"github.com/aykevl/tinygl"
	"tinygo.org/x/drivers/pixel"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
)

type BigText[T pixel.Color] struct {
	Name     string
	VBox     *tinygl.VBox[T]
	Header   *tinygl.Text[T]
	TextBox1 *tinygl.Text[T]
	TextBox2 *tinygl.Text[T]
}

// createWasmPage creates the screen when executing wasm on the badge.
func NewBigText[T pixel.Color](d *display.Device[T], hdr, txt1, txt2 string) *BigText[T] {
	if d == nil {
		return nil
	}

	header := d.Theme.NewText(hdr)
	header.SetBackground(pixel.NewColor[T](255, 0, 0))
	header.SetColor(pixel.NewColor[T](255, 255, 255))

	textbox1 := d.Theme.NewText(txt1)
	textbox1.SetAlign(tinygl.AlignCenter)
	textbox2 := d.Theme.NewText(txt2)
	textbox2.SetAlign(tinygl.AlignCenter)

	vbox := d.Theme.NewVBox(header, textbox1, textbox2)

	return &BigText[T]{
		Name:     hdr,
		VBox:     vbox,
		Header:   header,
		TextBox1: textbox1,
	}
}

func (bt *BigText[T]) Show(d *display.Device[T]) {
	d.Screen.SetChild(bt.VBox)
	d.Screen.Update()
}
