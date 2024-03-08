# WASMBadge

![1000004208](https://github.com/hybridgroup/mechanoid-examples/assets/5520/26809a01-ddcc-4bf2-a853-49a57fbddece)


This application is a conference badge programmed using WASM.

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

### Simulator

You need to install the Fyne cross-platform GUI toolkit to use the Mechanoid simulator.

https://github.com/fyne-io/fyne


```
$ mecha run -i wasman                                                                                                                                                        
Running using interpreter wasman                                                                                                                                             
Mechanoid engine starting...                                                                                                                                                 
Using interpreter wasman                                                                                                                                                     
Initializing engine...                                                                                                                                                       
Registering host modules...                                                                                                                                                  
Running WASM module mynameis.wasm                                                                                                                                            
Running module...                                                                                                                                                            
Running WASM module hithere.wasm                                                                                                                                             
Running module...                                                                                                                                                            
Mechanoid engine starting...                                                                                                                                                 
Using interpreter wasman                                                                                                                                                     
Initializing engine...                                                                                                                                                       
Registering host modules...                                                                                                                                                  
Running WASM module mynameis.wasm                                                                                                                                            
Running module...                                                                                                                                                            
Running WASM module hithere.wasm
Running module...
...
```