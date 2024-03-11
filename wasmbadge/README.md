# WASMBadge

![WASMBadge](../images/wasmbadge-pybadge.jpg)

This application is a conference badge programmed using WASM.

## How it works

The application can connect to any of the display supported in the `boards` package.

It embeds all of the WASM files in the `modules` directory right into the application itself.

When the application runs, it presents a list of all of the different programs on the display.

Use the buttons to choose one of the programs, and then press the "A" button to run it.

If you want to cycle thru the entire list, press the "START" button. The badge will run each of the WASM programs for 10 seconds before switching to the next one.

To get back to the home screen, press the "SELECT" button.

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