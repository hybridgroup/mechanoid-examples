package display

import (
	"machine"

	"image/color"

	"github.com/hybridgroup/mechanoid/engine"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var white = color.RGBA{255, 255, 255, 255}

type Device struct {
	eng     *engine.Engine
	display ssd1306.Device
}

// NewDevice creates a new display device.
func NewDevice(e *engine.Engine) *Device {
	return &Device{
		eng: e,
	}
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
	tinyfont.WriteLine(&d.display, &freemono.Bold9pt7b, int16(x), int16(y), msg, white)
	d.display.Display()
}
