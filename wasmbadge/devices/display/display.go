package display

import (
	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/basic"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/pixel"
)

type Device[T pixel.Color] struct {
	Display board.Displayer[T]
	Screen  *tinygl.Screen[T]
	Theme   *basic.Basic[T]
}

func NewDevice[T pixel.Color](disp board.Displayer[T]) Device[T] {
	return Device[T]{
		Display: disp,
	}
}

func (d *Device[T]) Init() error {
	// Determine size and scale of the screen.
	width, height := d.Display.Size()
	scalePercent := board.Display.PPI() * 100 / 120

	// Initialize the screen.
	buf := pixel.NewImage[T](int(width), int(height)/4)
	d.Screen = tinygl.NewScreen[T](d.Display, buf, board.Display.PPI())
	d.Theme = basic.NewTheme(style.NewScale(scalePercent), d.Screen)

	board.Display.SetBrightness(board.Display.MaxBrightness())

	return nil
}

func (d *Device[T]) Modules() wypes.Modules {
	return wypes.Modules{}
}
