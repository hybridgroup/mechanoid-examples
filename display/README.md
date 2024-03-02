# Display

Application that shows display integration using generics and the board package.

## How it works

The application can connect to any of the display supported in the `boards` package.

It then loads the `ping.wasm` program which is embedded into the application itself.

## How to run

### Build the WASM modules

```
$ mecha build                                                             
Building module ping
   code    data     bss |   flash     ram
      9       0       0 |       9       0
```

### Flash the board

PyBadge:

```
$ mecha flash -m pybadge
   code    data     bss |   flash     ram
 112188    2044    6696 |  114232    8740
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
Mechanoid engine starting...
Using interpreter...
Initializing engine...
Loading module...
Running module...
Ping 0
pong 1
Ping 1
pong 2
...
```

Gopher Badge

```
$ mecha flash -m gopher-badge
   code    data     bss |   flash     ram
 121452    2048    3228 |  123500    5276
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
Mechanoid engine starting...
Using interpreter...
Initializing engine...
Loading module...
Running module...
Ping 0
pong 1
Ping 1
pong 2
Ping 2
pong 3
...
```
