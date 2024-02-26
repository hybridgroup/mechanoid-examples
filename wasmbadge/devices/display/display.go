package display

import (
	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/basic"
	"tinygo.org/x/drivers/pixel"
)

type Device[T pixel.Color] struct {
	Display board.Displayer[T]
	Screen  *tinygl.Screen[T]
	Theme   *basic.Basic[T]
}

func NewDevice[T pixel.Color](disp board.Displayer[T]) Device[T] {
	// Determine size and scale of the screen.
	width, height := disp.Size()
	scalePercent := board.Display.PPI() * 100 / 120

	// Initialize the screen.
	buf := pixel.NewImage[T](int(width), int(height)/4)
	screen := tinygl.NewScreen[T](disp, buf, board.Display.PPI())
	theme := basic.NewTheme(style.NewScale(scalePercent), screen)

	board.Display.SetBrightness(board.Display.MaxBrightness())

	return Device[T]{
		Display: disp,
		Screen:  screen,
		Theme:   theme,
	}
}
