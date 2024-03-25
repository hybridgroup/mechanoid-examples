# pingzig

WASM unknown module written in Zig

## How to install

- Install Zig.

## How to build module

```bash
cd modules/pingzig
zig build-lib -rdynamic -dynamic -target wasm32-freestanding -OReleaseSmall --stack 4096 --import-memory --initial-memory=65536 --max-memory=65536 ping.zig
cd ../..
cp ./modules/pingzig/ping.wasm ./modules/pingzig.wasm
```
