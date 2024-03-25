# pingrs

WASM unknown module written in Rust

## How to install

- Install Rust.

- Install the `wasm32-unknown-unknown` target.

```bash
rustup target add wasm32-unknown-unknown
```

## How to build module

```bash
cd modules/pingrs
cargo build --target wasm32-unknown-unknown --release
cd ../..
cp ./modules/pingrs/target/wasm32-unknown-unknown/release/pingrs.wasm ./modules/
```
