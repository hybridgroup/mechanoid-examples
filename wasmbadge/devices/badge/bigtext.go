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
func NewBigText[T pixel.Color](d *display.Device[T]) *BigText[T] {
	if d == nil {
		return nil
	}

	header := d.Theme.NewText("Mechanoid")
	header.SetBackground(pixel.NewColor[T](255, 0, 0))
	header.SetColor(pixel.NewColor[T](255, 255, 255))

	textbox1 := d.Theme.NewText("")
	textbox1.SetAlign(tinygl.AlignCenter)
	textbox2 := d.Theme.NewText("")
	textbox2.SetAlign(tinygl.AlignCenter)

	vbox := d.Theme.NewVBox(header, textbox1, textbox2)

	return &BigText[T]{
		Name:     "BigText",
		VBox:     vbox,
		Header:   header,
		TextBox1: textbox1,
		TextBox2: textbox2,
	}
}

func (bt *BigText[T]) Show(d *display.Device[T]) {
	d.Screen.SetChild(bt.VBox)
	d.Screen.Update()
}

func (bt *BigText[T]) Heading(s string) {
	bt.Header.SetText(s)
}

func (bt *BigText[T]) SetText1(s string) {
	bt.TextBox1.SetText(s)
}

func (bt *BigText[T]) SetText2(s string) {
	bt.TextBox2.SetText(s)
}
