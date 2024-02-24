# Blink

Example that loads a WASM module `blink.wasm` that can blink an LED.

## How to run

Create a new project based on this example:

```
mecha new example.com/modules/blinky github.com/hybridgroup/mechanoid-examples/blinky
```

Now install and build the needed WASM module:

```
mecha new module example.com/modules/blink github.com/hybridgroup/mechanoid-examples/modules/blink
mecha build
```

Now you can build and flash your board:

```
mecha flash pybadge -monitor
```
