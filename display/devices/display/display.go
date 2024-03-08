package display

import (
	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/basic"
	"github.com/hybridgroup/mechanoid/convert"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/pixel"
)

var pongCount int

type Device[T pixel.Color] struct {
	Display  board.Displayer[T]
	Screen   *tinygl.Screen[T]
	Theme    *basic.Basic[T]
	VBox     *tinygl.VBox[T]
	Header   *tinygl.Text[T]
	PingText *tinygl.Text[T]
	PongText *tinygl.Text[T]
}

// NewDevice creates a new display device.
func NewDevice[T pixel.Color](display board.Displayer[T]) *Device[T] {
	return &Device[T]{
		Display: display,
	}
}

func (d *Device[T]) Modules() wypes.Modules {
	return wypes.Modules{ //}
		"hosted": wypes.Module{
			"pong": wypes.H0(d.pongFunc),
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
	d.Header = d.Theme.NewText("Hello, Mechanoid")
	d.PingText = d.Theme.NewText("waiting...")
	d.PongText = d.Theme.NewText("waiting...")
	d.VBox = d.Theme.NewVBox(d.Header, d.PingText, d.PongText)

	d.Screen.SetChild(d.VBox)
	d.Screen.Update()
	board.Display.SetBrightness(board.Display.MaxBrightness())

	return nil
}

func (d *Device[T]) ShowPing(count int) {
	msg := "Ping: " + convert.IntToString(count)
	d.PingText.SetText(msg)
	d.Screen.Update()
}

func (d *Device[T]) ShowPong(count int) {
	msg := "Pong: " + convert.IntToString(count)
	d.PongText.SetText(msg)
	d.Screen.Update()
}

func (d *Device[T]) pongFunc() wypes.Void {
	pongCount++
	println("Pong", pongCount)
	d.ShowPong(pongCount)
	return wypes.Void{}
}
