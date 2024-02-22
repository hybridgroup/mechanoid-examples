# Ping

Exports a `ping()` function, that immediately calls the host's `pong()` function.

## Building

```
tinygo build -size short -o ./modules/ping/ping.wasm -target ./modules/ping/wasm-unknown.json -no-debug ./modules/ping
```
