package main

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

const (
	TakeoffNormal = iota
	TakeoffThrow
)

const (
	LandingNormal = iota
	LandingHand
)

const (
	FlipFront = iota
	FlipLeft
	FlipBack
	FlipRight
	FlipForwardLeft
	FlipBackLeft
	FlipBackRight
	FlipForwardRight
)

const defaultSpeed = 30

//go:wasmimport drone control
func droneControl(direction, speed uint32)

//go:wasmimport drone takeoff
func droneTakeoff(kind uint32)

//go:wasmimport drone land
func droneLand(kind uint32)

//go:wasmimport drone flip
func droneFlip(kind uint32)
