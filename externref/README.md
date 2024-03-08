# Externalref

Here is an example of an application built using Mechanoid that uses host external references.

For more information, see https://webassembly.github.io/gc/core/syntax/types.html#reference-types

## How it works

The application creates a new module.

It then loads the `hollaback.wasm` program which is embedded into the application itself.

That module is able to obtain an externref to the host instance, and then use it to call the associated host method.

## How to run

### PyBadge

```
$ mecha flash -i wazero -m pybadge
Building module hollaback  
Done.            
   code    data     bss |   flash     ram
    112      33   65503 |     145   65536
Application built. Now flashing...
   code    data     bss |   flash     ram
 331152   66348    7112 |  397500   73460
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
Mechanoid engine starting...
Using interpreter wazero   
Initializing engine...
Initializing interpreter...
Initializing devices...    
Defining host function...
Loading WASM module...
Running module...          
Calling start...
newGreeter msg is Hello, WebAssembly!                                                 
got value: 1
Calling update...                                                                     
got value: 1
hello msg is From Mechanoid
Calling update...
got value: 1
hello msg is From Mechanoid
Calling update...
got value: 1
...
```
