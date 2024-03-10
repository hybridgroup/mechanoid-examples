package display

import (
	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/basic"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/pixel"
)

var pongCount int

type Device[T pixel.Color] struct {
	Display      board.Displayer[T]
	Screen       *tinygl.Screen[T]
	Theme        *basic.Basic[T]
	VBox         *tinygl.VBox[T]
	Header       *tinygl.Text[T]
	Message1Text *tinygl.Text[T]
	Message2Text *tinygl.Text[T]
}

// NewDevice creates a new display device.
func NewDevice[T pixel.Color](display board.Displayer[T]) *Device[T] {
	return &Device[T]{
		Display: display,
	}
}

func (d *Device[T]) Modules() wypes.Modules {
	return wypes.Modules{
		"display": wypes.Module{
			"heading":  wypes.H1(d.heading),
			"message1": wypes.H1(d.message1),
			"message2": wypes.H1(d.message2),
		},
	}
}

func (d *Device[T]) Init() error {
	width, height := d.Display.Size()
	scalePercent := board.Display.PPI() * 100 / 120

	// Initialize the screen.
	buf := pixel.NewImage[T](int(width), int(height)/4)
	d.Screen = tinygl.NewScreen[T](d.Display, buf, board.Display.PPI())
	d.Theme = basic.NewTheme(style.NewScale(scalePercent), d.Screen)
	d.Header = d.Theme.NewText("Mechanoid")
	d.Message1Text = d.Theme.NewText("")
	d.Message2Text = d.Theme.NewText("")
	d.VBox = d.Theme.NewVBox(d.Header, d.Message1Text, d.Message2Text)

	d.Screen.SetChild(d.VBox)
	d.Screen.Update()
	board.Display.SetBrightness(board.Display.MaxBrightness())

	return nil
}

func (d *Device[T]) Heading(msg string) {
	d.Header.SetText(msg)
	d.Screen.Update()
}

func (d *Device[T]) Message1(msg string) {
	d.Message1Text.SetText(msg)
	d.Screen.Update()
}

func (d *Device[T]) Message2(msg string) {
	d.Message2Text.SetText(msg)
	d.Screen.Update()
}

func (d *Device[T]) heading(msg wypes.String) wypes.Void {
	d.Heading(msg.Unwrap())
	return wypes.Void{}
}

func (d *Device[T]) message1(msg wypes.String) wypes.Void {
	d.Message1(msg.Unwrap())
	return wypes.Void{}
}

func (d *Device[T]) message2(msg wypes.String) wypes.Void {
	d.Message2(msg.Unwrap())
	return wypes.Void{}
}
