//go:build tinygo

package main

import (
	_ "embed"
)

//go:embed tinygo-logo.qoi
var data string

var (
	img *Image
)

//go:export start
func start() {
	img = NewImage(data)
	img.Show()
}

//go:export update
func update() {
}

func main() {}
