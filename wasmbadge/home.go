package main

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style/basic"
	"tinygo.org/x/drivers/pixel"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
)

type HomePage[T pixel.Color] struct {
	Name    string
	VBox    *tinygl.VerticalScrollBox[T]
	Header  *tinygl.Text[T]
	ListBox *basic.ListBox[T]
}

var (
	menuChoices = make([]string, 0, 10)
)

// createHome creates the home screen for the badge.
func createHome[T pixel.Color](d *display.Device[T]) *HomePage[T] {
	// Create badge homescreen.
	header := d.Theme.NewText("WASM Badge")
	header.SetBackground(pixel.NewColor[T](255, 0, 0))
	header.SetColor(pixel.NewColor[T](255, 255, 255))
	listbox := d.Theme.NewListBox(menuChoices)
	listbox.SetGrowable(0, 1) // listbox fills the rest of the screen
	listbox.Select(0)         // focus the first element
	home := tinygl.NewVerticalScrollBox[T](header, listbox, nil)
	return &HomePage[T]{
		Name:    "home",
		VBox:    home,
		Header:  header,
		ListBox: listbox,
	}
}

func (p *HomePage[T]) Show(d *display.Device[T]) {
	d.Screen.SetChild(p.VBox)
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
