//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

var (
	// buf is a buffer that is used to pass messages to the host instance.
	buf [64]byte
)

//go:export button_select
func buttonSelect(shift int32) {
	msg := "Flip!"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message1(ptr, sz)

	kind := FlipFront
	if shift == 1 {
		kind = FlipBack
	}
	droneFlip(uint32(kind))
}

//go:export button_start
func buttonStart(shift int32) {
	msg := "Takeoff"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message1(ptr, sz)

	kind := TakeoffNormal
	if shift == 1 {
		kind = TakeoffThrow
	}
	droneTakeoff(uint32(kind))
}

//go:export button_b
func buttonB(shift int32) {
	msg := "Landing"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message1(ptr, sz)

	kind := LandingNormal
	if shift == 1 {
		kind = LandingHand
	}
	droneLand(uint32(kind))
}

//go:export button_up
func buttonUp(shift int32) {
	msg := "UP"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])
	message1(ptr, sz)

	direction := DirectionForward
	if shift == 1 {
		direction = DirectionUp
	}

	droneControl(uint32(direction), defaultSpeed)
}

//go:export button_down
func buttonDown(shift int32) {
	msg := "DOWN"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message1(ptr, sz)

	direction := DirectionBackward
	if shift == 1 {
		direction = DirectionDown
	}

	droneControl(uint32(direction), defaultSpeed)
}

//go:export button_left
func buttonLeft(shift int32) {
	msg := "LEFT"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message1(ptr, sz)

	direction := DirectionLeft
	if shift == 1 {
		direction = DirectionTurnLeft
	}

	droneControl(uint32(direction), defaultSpeed)
}

//go:export button_right
func buttonRight(shift int32) {
	msg := "RIGHT"
	copy(buf[:], msg)
	ptr, sz := convert.BytesToWasmPtr(buf[:len(msg)])

	message1(ptr, sz)

	direction := DirectionRight
	if shift == 1 {
		direction = DirectionTurnRight
	}

	droneControl(uint32(direction), defaultSpeed)
}

func main() {}
