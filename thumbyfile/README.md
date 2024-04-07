# Filestore

Application that demonstrates how to use the onboard Flash storage on the hardware device to save/load/run external WASM modules via a Command line interface directly on the device itself.

## How to run

### Flash the board

```bash
$ mecha flash -m pybadge                                    
Building TinyGo module ping                          
Done.                                               
code    data     bss |   flash     ram
   9       0       0 |       9       0             
Building Rust module pingrs                                
Done.                          
warning: unstable feature specified for `-Ctarget-feature`: `atomics`                                                
  | 
  = note: this feature is not stably supported; its behavior can change in the future
warning: unstable feature specified for `-Ctarget-feature`: `bulk-memory`
  | 
  = note: this feature is not stably supported; its behavior can change in the future 
warning: `pingrs` (lib) generated 2 warnings
Finished release [optimized] target(s) in 0.00s
Building Zig module pingzig
Done.                       
Application built. Now flashing...
code    data     bss |   flash     ram                     
342556   16812    7224 |  359368   2403
Connected to /dev/ttyACM0. Press Ctrl-C to exit.

==>
```

You should see the `==>` prompt. See "How to use" below.

## How to use

You should see a `==>` prompt. Try the `lsblk` command to see the Flash storage information:

```bash
==> lsblk
-------------------------------------
 Device Information:
-------------------------------------
 flash data start: 0x00024000
 flash data end:   0x00080000
-------------------------------------
```

This the the available Flash memory on your board in the extra space not being used by your program.

Try the `ls` command.

```bash
==> ls

-------------------------------------
 File Store:
-------------------------------------

-------------------------------------
```

You do not yet have any WASM files in the Flash storage. Let's put one on the device using the `save` command.

The easiest way to do this is the included `savefile` program. Press `CTRL-C` to return to your shell, then run the following command (substitute the correct port name for `/dev/ttyACM0` as needed):

```bash
cd ./filestore

go run ./savefile ./modules/ping.wasm /dev/ttyACM0
```

Now connect again to the board, and now you should see the file listed using the `ls` command:

```bash
$ mecha monitor
Connected to /dev/ttyACM0. Press Ctrl-C to exit.

==> ls

-------------------------------------
 File Store:
-------------------------------------
370 ping.wasm

-------------------------------------
```

You can now load the module:

```bash
==> load ping.wasm
loading ping.wasm
module loaded.
```

And then start it running:

```bash
==> run
module running.
```

Use the `ping` command:

```bash
==> ping 3
Ping...
pong
Ping...
pong
Ping...
pong
```

Use the `halt` command to stop the module. 

```bash
==> halt                                                                              
halting...
module halted.
```

You can load another module now. Let's try one written using Rust.

First, transfer the compiled pingrs.wasm module to the board. Press `CTRL-C` to return to your shell, then run the following command:

```bash
go run ./savefile ./modules/pingrs.wasm /dev/ttyACM0
```

Now connect again to the board, and now you should see the file listed using the `ls` command, alongside the previously saved file:

```bash
$ mecha monitor
Connected to /dev/ttyACM0. Press Ctrl-C to exit.

==> ls

-------------------------------------
 File Store:
-------------------------------------
370 ping.wasm
324 pingrs.wasm
-------------------------------------
```

You can now load the Rust module:

```bash
==> load pingrs.wasm
loading pingrs.wasm
module loaded.
```

start it running:

```bash
==> run
module running.
```

And use the `ping` command to call the Rust module:

```bash
==> ping 3
Ping...
pong
Ping...
pong
Ping...
pong
```
