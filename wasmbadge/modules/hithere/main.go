//go:build tinygo

package main

import (
	"github.com/hybridgroup/mechanoid/convert"
)

const msg = "Hello, WebAssembly!"

var (
	bgtext  *BigText
	counter int
	buf     [64]byte
)

//go:export start
func start() {
	bgtext = NewBigText(msg)
}

//go:export update
func update() {
	start, end := 0, len(msg)
	copy(buf[:], msg)
	bgtext.SetText1(buf[start:end])

	m := "Count: "
	start, end = end, end+len(m)
	copy(buf[start:end], m)

	counter++
	s := convert.IntToString(counter)
	st2 := end
	end = end + len(s)

	copy(buf[st2:end], s)
	bgtext.SetText2(buf[start:end])

	bgtext.Show()
}

func main() {}
