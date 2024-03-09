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
	Display     board.Displayer[T]
	Screen      *tinygl.Screen[T]
	Theme       *basic.Basic[T]
	VBox        *tinygl.VBox[T]
	Header      *tinygl.Text[T]
	MessageText *tinygl.Text[T]
}

// NewDevice creates a new display device.
func NewDevice[T pixel.Color](display board.Displayer[T]) *Device[T] {
	return &Device[T]{
		Display: display,
	}
}

func (d *Device[T]) Modules() wypes.Modules {
	return wypes.Modules{ //}
		"display": wypes.Module{
			"message": wypes.H1(d.showMessage),
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
	d.Header = d.Theme.NewText("Mechanoid Buttons")
	d.MessageText = d.Theme.NewText("waiting...")
	d.VBox = d.Theme.NewVBox(d.Header, d.MessageText)

	d.Screen.SetChild(d.VBox)
	d.Screen.Update()
	board.Display.SetBrightness(board.Display.MaxBrightness())

	return nil
}

func (d *Device[T]) ShowMessage(msg string) {
	d.MessageText.SetText(msg)
	d.Screen.Update()
}

func (d *Device[T]) showMessage(msg wypes.String) wypes.Void {
	d.ShowMessage(msg.Unwrap())
	return wypes.Void{}
}
