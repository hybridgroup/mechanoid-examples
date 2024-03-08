package display

import (
	"machine"

	"image/color"

	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

var white = color.RGBA{255, 255, 255, 255}

type Device struct {
	display ssd1306.Device
}

func (d *Device) Modules() wypes.Modules {
	return wypes.Modules{}
}

func (d *Device) Init() error {
	machine.SPI0.Configure(machine.SPIConfig{})
	display := ssd1306.NewSPI(machine.SPI0, machine.THUMBY_DC_PIN, machine.THUMBY_RESET_PIN, machine.THUMBY_CS_PIN)
	display.Configure(ssd1306.Config{
		Width:     72,
		Height:    40,
		ResetCol:  ssd1306.ResetValue{28, 99},
		ResetPage: ssd1306.ResetValue{0, 5},
	})

	display.ClearDisplay()
	d.display = display

	return nil
}

func (d *Device) Clear() {
	d.display.ClearDisplay()
}

func (d *Device) ShowMessage(x, y int, msg string) {
	tinyfont.WriteLine(&d.display, &proggy.TinySZ8pt7b, int16(x), int16(y), msg, white)
	d.display.Display()
}
