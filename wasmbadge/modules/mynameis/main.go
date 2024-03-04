//go:build tinygo

package main

var (
	bgtext *BigText
)

const msg = "TinyGo"

//go:export start
func start() {
	bgtext = NewBigText(msg)
}

//go:export update
func update() {
	// TODO: something with the ui
}

func main() {}
