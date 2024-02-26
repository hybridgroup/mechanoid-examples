package badge

import (
	"github.com/hybridgroup/mechanoid-examples/wasmbadge/devices/display"
	"github.com/hybridgroup/mechanoid/engine"
	"tinygo.org/x/drivers/pixel"
)

type Badge[T pixel.Color] struct {
	Engine  *engine.Engine
	Display *display.Device[T]
}

const moduleName = "badge"

func NewDevice[T pixel.Color](e *engine.Engine) *Badge[T] {
	return &Badge[T]{
		Engine: e,
	}
}

func (b *Badge[T]) Init() error {
	if b.Engine == nil {
		return engine.ErrInvalidEngine
	}

	if err := b.Engine.Interpreter.DefineFunc(moduleName, "new_big_text", b.newBigText); err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func (b *Badge[T]) UseDisplay(d *display.Device[T]) error {
	b.Display = d
	return nil
}

func (b *Badge[T]) newBigText(ptr uint32, sz uint32) uint32 {
	msg, err := b.Engine.Interpreter.MemoryData(ptr, sz)
	if err != nil {
		println(err.Error())
		return 0
	}

	// create the badge UI element
	bt := NewBigText[T](b.Display, "WASM Badge", string(msg), "")
	if bt == nil {
		return 0
	}

	bt.Show(b.Display)
	return sz
}
