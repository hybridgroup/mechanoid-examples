![Mechanoid logo](https://mechanoid.io/images/logo-blue.png)

# Mechanoid Examples

This repo contains example applications written using Mechanoid (https://mechanoid.io)

See the README files inside each for details on what they do, how they work and the **recommended runtime interpreter** (wazero, wasman,...), as there are some differences among them and the example might not run correctly. 

## Blinky

![Blinky](./images/blinky-pybadge.jpg)

Application that loads an embedded WASM program that blinks an LED on the hardware.

## Buttons

![Buttons](./images/buttons-gopher-badge.jpg)

Application that loads an embedded WASM program and sends it events from pressing the buttons. The WASM programs then displays messages on the small screen on the hardware device.

## Display

![Display](./images/display-pybadge.jpg)

Application that loads an embedded WASM program and displays the interaction on the small screen on the hardware device.

## Externref

Application built using Mechanoid that uses host external references.

## Filestore

Application that demonstrates how to use the onboard Flash storage on the hardware device to save/load/run external WASM modules via a Command line interface directly on the device itself.

Also shows how to use Mechanoid with WASM modules written using TinyGo, Rust, or Zig.

## Simple

Example of a simple "ping" application built using Mechanoid.

## Thumby

![Thumby](./images/thumby.jpg)

This is an example of an application built using Mechanoid specifically for the Thumby "itty-bitty game system". 

## ThumbyFile

![Thumby](./images/thumby.jpg)

Application that demonstrates how to use the onboard Flash storage on the Thumby device to save/load/run external WASM modules via a Command line interface directly on the device itself, along with display support so you can see what it happening on the tiny display.

Also shows how to use Mechanoid with WASM modules written using TinyGo, Rust, or Zig.

## WASMBadge

![WASMBadge](./images/wasmbadge-pybadge.jpg)

This application is a conference badge programmed using WASM.

## WASMDrone

![WASMDrone](./images/wasmdrone-pybadge-tello.jpg)

This application lets you write a WebAssembly program that runs on a hardware device with connected wireless to fly a DJI Tello drone.
