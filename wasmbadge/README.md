# WASMBadge

![1000004208](https://github.com/hybridgroup/mechanoid-examples/assets/5520/26809a01-ddcc-4bf2-a853-49a57fbddece)


Conference badge programmed using WASM.

## How it works

The application can connect to any of the display supported in the `boards` package.

It then loads the `ping.wasm` program which is embedded into the application itself.

## How to run

### PyBadge

```
$ mecha flash -i wasman pybadge
Building module hithere
Done.
   code    data     bss |   flash     ram
    576      31    4097 |     607    4128
Building module mynameis
Done.
   code    data     bss |   flash     ram
     15       6    4096 |      21    4102
Application built. Now flashing...
   code    data     bss |   flash     ram
 131380    1740    6792 |  133120    8532
```
