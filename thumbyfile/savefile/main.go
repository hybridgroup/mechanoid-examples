package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"go.bug.st/serial"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s path/to/xxx.wasm port", os.Args[0])
		os.Exit(0)
	}
	err := run(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

func run(target, port string) error {
	p, err := serial.Open(port, &serial.Mode{})
	if err != nil {
		return err
	}
	defer p.Close()
	p.SetReadTimeout(1 * time.Second)

	b, err := os.ReadFile(target)
	if err != nil {
		return err
	}

	fmt.Fprintf(p, "\r\nsave %s %d\r\n", filepath.Base(target), len(b))

	buf := make([]byte, 1024)
	for {
		n, err := p.Read(buf)
		if err != nil || n == 0 {
			break
		}
	}

	_, err = p.Write(b)
	if err != nil {
		return err
	}

	return nil
}
