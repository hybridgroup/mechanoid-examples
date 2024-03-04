package badge

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/image"
	"tinygo.org/x/drivers/pixel"

	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
)

type Image[T pixel.Color] struct {
	Name   string
	VBox   *tinygl.VBox[T]
	Header *tinygl.Text[T]
	Image  *tinygl.Image[T]
}

func NewImage[T pixel.Color](d *display.Device[T], hdr string, img image.Image[T]) *Image[T] {
	if d == nil {
		return nil
	}

	header := d.Theme.NewText(hdr)
	header.SetBackground(pixel.NewColor[T](255, 0, 0))
	header.SetColor(pixel.NewColor[T](255, 255, 255))

	image1 := tinygl.NewImage(img)

	vbox := d.Theme.NewVBox(header, image1)

	return &Image[T]{
		Name:   hdr,
		VBox:   vbox,
		Header: header,
		Image:  image1,
	}
}

func (i *Image[T]) Show(d *display.Device[T]) {
	d.Screen.SetChild(i.VBox)
	d.Screen.Update()
}
