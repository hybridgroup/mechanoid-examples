# WASMBadge

![1000004208](https://github.com/hybridgroup/mechanoid-examples/assets/5520/26809a01-ddcc-4bf2-a853-49a57fbddece)


Conference badge programmed using WASM.

## How it works

The application can connect to any of the display supported in the `boards` package.

It then loads the `ping.wasm` program which is embedded into the application itself.

## How to run

First compile all of the built-in WASM modules:

```
mecha build
```

Now flash the application on to your badge.

### PyBadge

```
mecha flash pybadge
```

### Gopher Badge

```
mecha flash gopher-badge
```
