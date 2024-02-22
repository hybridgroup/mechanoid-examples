//go:build tinygo

package main

import (
	"machine"
)

var (
	led = machine.Pin(23)
	on  bool
)

//go:export setup
func setup() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

//go:export loop
func loop() {
	on = !on
	led.Set(on)
	on = led.Get()
}

func main() {}
