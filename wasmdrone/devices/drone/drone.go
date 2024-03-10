package drone

import (
	"errors"
	"time"

	"github.com/hybridgroup/mechanoid"
	tello "github.com/hybridgroup/tinygo-tello"
	"github.com/orsinium-labs/wypes"
	"tinygo.org/x/drivers/netlink"
	"tinygo.org/x/drivers/netlink/probe"
)

const (
	DirectionNone = iota
	DirectionForward
	DirectionBackward
	DirectionLeft
	DirectionRight
	DirectionUp
	DirectionDown
	DirectionTurnLeft
	DirectionTurnRight
)

var errNoSSID = errors.New("no SSID provided")

type Device struct {
	drone      *tello.Tello
	ssid, pass string
	connected  bool
	Direction  int
	Speed      int
}

func NewDevice(ssid, pass string) *Device {
	return &Device{
		ssid: ssid,
		pass: pass,
	}
}

func (d *Device) Modules() wypes.Modules {
	return wypes.Modules{
		"drone": wypes.Module{
			"control": wypes.H2(d.control),
			"takeoff": wypes.H1(d.takeoff),
			"land":    wypes.H1(d.land),
			"flip":    wypes.H1(d.flip),
		},
	}
}

func (d *Device) control(direction, speed wypes.UInt32) wypes.Void {
	d.Direction = int(direction.Unwrap())
	d.Speed = int(speed.Unwrap())

	return wypes.Void{}
}

func (d *Device) takeoff(kind wypes.UInt32) wypes.Void {
	// TODO: simple debounce
	switch kind.Unwrap() {
	case 0:
		d.TakeOff()
	case 1:
		d.ThrowTakeOff()
	}
	return wypes.Void{}
}

func (d *Device) land(kind wypes.UInt32) wypes.Void {
	// TODO: simple debounce
	switch kind.Unwrap() {
	case 0:
		d.Land()
	case 1:
		d.PalmLand()
	}
	return wypes.Void{}
}

func (d *Device) flip(kind wypes.UInt32) wypes.Void {
	// TODO: simple debounce
	d.Flip(int(kind.Unwrap()))
	return wypes.Void{}
}

func (d *Device) Init() error {
	if d.ssid == "" {
		return errNoSSID
	}

	mechanoid.Debug("connecting to drone")
	link, _ := probe.Probe()
	err := link.NetConnect(&netlink.ConnectParams{
		Ssid:       d.ssid,
		Passphrase: d.pass,
	})

	if err != nil {
		return err
	}

	d.drone = tello.New("8888")

	return nil
}

func (d *Device) Start() error {
	if err := d.drone.Start(); err != nil {
		return err
	}

	d.connected = true
	return nil
}

func (d *Device) Control() {
	for {
		if !d.connected {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		switch d.Direction {
		case DirectionForward:
			d.drone.Forward(d.Speed)
		case DirectionBackward:
			d.drone.Backward(d.Speed)
		default:
			d.drone.Forward(0)
		}

		switch d.Direction {
		case DirectionLeft:
			d.drone.Left(d.Speed)
		case DirectionRight:
			d.drone.Right(d.Speed)
		default:
			d.drone.Right(0)
		}

		switch d.Direction {
		case DirectionUp:
			d.drone.Up(d.Speed)
		case DirectionDown:
			d.drone.Down(d.Speed)
		default:
			d.drone.Up(0)
		}

		switch d.Direction {
		case DirectionTurnLeft:
			d.drone.CounterClockwise(d.Speed)
		case DirectionTurnRight:
			d.drone.Clockwise(d.Speed)
		default:
			d.drone.Clockwise(0)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (d *Device) TakeOff() error {
	return d.drone.TakeOff()
}

func (d *Device) ThrowTakeOff() error {
	return d.drone.ThrowTakeOff()
}

func (d *Device) Land() error {
	return d.drone.Land()
}

func (d *Device) PalmLand() error {
	return d.drone.PalmLand()
}

func (d *Device) Flip(kind int) error {
	return d.drone.Flip(tello.FlipType(kind))
}
