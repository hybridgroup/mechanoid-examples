# Display

Application that shows display integration using generics and the board package.

## How it works

The application can connect to any of the display supported in the `boards` package.

It then loads the `ping.wasm` program which is embedded into the application itself.

## How to run

### PyBadge

```
tinygo flash -size short -target pybadge ./display
```

### Gopher Badge

```
tinygo flash -size short -target gopher-badge -stack-size=4kb ./display
```
