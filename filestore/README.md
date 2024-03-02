# Filestore

Application that has a Command line interface to save/load/run WASM modules using the onboard Flash storage.


## How to run

### Flash the board

#### PyBadge

```
$ mecha flash -m pybadge
   code    data     bss |   flash     ram
 130368    2148    6888 |  132516    9036

Connected to /dev/ttyACM0. Press Ctrl-C to exit.

==> 

```

You should see the `==>` prompt. See "How to use" below.

#### Gopher Badge

```
$ mecha flash -m gopher-badge
   code    data     bss |   flash     ram
 139772       4    3632 |  139776    3636
Connected to /dev/ttyACM0. Press Ctrl-C to exit.

==>

```

You should see the `==>` prompt. See "How to use" below.


## How to use

You should see a `==>` prompt. Try the `lsblk` command to see the Flash storage information:

```
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

```
==> ls         
                                    
-------------------------------------                                    
 File Store:  
-------------------------------------
                                    
-------------------------------------
```

You do not yet have any WASM files in the Flash storage. Let's put one on the device using the `save` command.

The easiest way to do this is the included `savefile` program. Press `CTRL-C` to return to your shell, then run the following command (substitute the correct port name for `/dev/ttyACM0` as needed):

```
cd ./filestore

go run ./savefile ./ping.wasm /dev/ttyACM0
```

Now connect again to the board, and now you should see the file listed using the `ls` command:

```
$ tinygo monitor
Connected to /dev/ttyACM0. Press Ctrl-C to exit.

==> ls

-------------------------------------
 File Store:  
-------------------------------------
370 ping.wasm

-------------------------------------
```

You can now load the module:

```
==> load ping.wasm
load: ping.wasm
module loaded
```

And then start it running:

```
==> run
starting...
building index space
initializing memory
initializing functions
initializing globals
running start func
running.
```

Use the `ping` command:

```
==> ping 3
Ping...
pong
Ping...
pong
Ping...
pong
```

Use the `halt` command to stop the module. You can load another module now.
