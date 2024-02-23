//go:build tinygo

package main

//go:export hola
func hola(msg string) uint32

const msg = "Hello, WebAssembly!"

//go:export hello
func hello() {
	hola(msg)
}

func main() {}
