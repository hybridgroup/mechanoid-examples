# WASMBadge

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

## How to run

Create a new project based on this example:

```
mecha new example.com/modules/display github.com/hybridgroup/mechanoid-templates/display
```

Now install and build the needed WASM module:

```
cd display
mecha new module example.com/display/modules/blink github.com/hybridgroup/mechanoid-templates/modules/blink
mecha build
```

Now you can build and flash your board:

- PyBadge

```
mecha flash pybadge
```

- Gopher Badge

```
mecha flash gopher-badge
```
