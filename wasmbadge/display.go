package main

import (
	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/basic"
	"tinygo.org/x/drivers/pixel"
)

type DisplayDevice[T pixel.Color] struct {
	Display board.Displayer[T]
	Screen  *tinygl.Screen[T]
	Theme   *basic.Basic[T]
}

func NewDisplayDevice[T pixel.Color](disp board.Displayer[T]) DisplayDevice[T] {
	// Determine size and scale of the screen.
	width, height := disp.Size()
	scalePercent := board.Display.PPI() * 100 / 120

	// Initialize the screen.
	buf := pixel.NewImage[T](int(width), int(height)/4)
	screen := tinygl.NewScreen[T](disp, buf, board.Display.PPI())
	theme := basic.NewTheme(style.NewScale(scalePercent), screen)

	board.Display.SetBrightness(board.Display.MaxBrightness())

	return DisplayDevice[T]{
		Display: disp,
		Screen:  screen,
		Theme:   theme,
	}
}

type Page[T pixel.Color] struct {
	Name    string
	VBox    *tinygl.VerticalScrollBox[T]
	Header  *tinygl.Text[T]
	ListBox *basic.ListBox[T]
}

var (
	homePage    any
	menuChoices = make([]string, 0, 10)
)

// createHome creates the home screen for the badge.
func (d *DisplayDevice[T]) createHome() {
	// Create badge homescreen.
	header := d.Theme.NewText("WASM Badge")
	header.SetBackground(pixel.NewColor[T](255, 0, 0))
	header.SetColor(pixel.NewColor[T](255, 255, 255))
	listbox := d.Theme.NewListBox(menuChoices)
	listbox.SetGrowable(0, 1) // listbox fills the rest of the screen
	listbox.Select(0)         // focus the first element
	home := tinygl.NewVerticalScrollBox[T](header, listbox, nil)
	homePage = &Page[T]{
		Name:    "home",
		VBox:    home,
		Header:  header,
		ListBox: listbox,
	}
}

func (d *DisplayDevice[T]) showHome() {
	d.Screen.SetChild(homePage.(*Page[T]).VBox)
	d.Screen.Update()
}

func loadMenuChoices() error {
	modules, err := modules.ReadDir("modules")
	if err != nil {
		return err
	}

	for _, module := range modules {
		if module.IsDir() {
			continue
		}
		menuChoices = append(menuChoices, module.Name())
	}
	return nil
}

var (
	wasmPage any
)

type WASMPage[T pixel.Color] struct {
	Name    string
	VBox    *tinygl.VBox[T]
	Header  *tinygl.Text[T]
	TextBox *tinygl.Text[T]
}

// createWasmPage creates the screen when executing wasm on the badge.
func (d *DisplayDevice[T]) createWasmPage(module string) {
	header := d.Theme.NewText(module)
	header.SetBackground(pixel.NewColor[T](255, 0, 0))
	header.SetColor(pixel.NewColor[T](255, 255, 255))

	textbox := d.Theme.NewText("Running " + module)
	textbox.SetAlign(tinygl.AlignCenter)
	wasmbox := d.Theme.NewVBox(header, textbox)
	wasmPage = &WASMPage[T]{
		Name:    module,
		VBox:    wasmbox,
		Header:  header,
		TextBox: textbox,
	}
}

func (d *DisplayDevice[T]) showWasmPage() {
	d.Screen.SetChild(wasmPage.(*WASMPage[T]).VBox)
	d.Screen.Update()
}

func (d *DisplayDevice[T]) outputToWasmPage(s string) {
	wasmPage.(*WASMPage[T]).TextBox.SetText(s)
	d.Screen.Update()
}
