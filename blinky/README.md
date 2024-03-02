# Blink

Example that loads a WASM module `blink.wasm` that can blink an LED.

## How to run

### Build the WASM modules

```
$ mecha build
Building module blink
   code    data     bss |   flash     ram
     68       0       0 |      68       0
```

### Flash the board

```
$ mecha flash -m pybadge
   code    data     bss |   flash     ram
 103748    2316    6680 |  106064    8996
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
Mechanoid engine starting...
Using interpreter...
Initializing engine...
Loading module...
Running module...
Calling setup...
Calling loop...
Calling loop...
Calling loop...
```
