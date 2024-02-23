(module
  (type (;0;) (func))
  (type (;1;) (func (param i32 i32) (result i32)))
  (import "env" "hola" (func $hola (type 1)))
  (import "env" "memory" (memory (;0;) 1 1))
  (func $__wasm_call_dtors (type 0))
  (func $hello (type 0)
    i32.const 4096
    i32.const 19
    call $hola
    drop)
  (export "_initialize" (func $__wasm_call_dtors))
  (export "hello" (func $hello))
  (data (;0;) (i32.const 4096) "Hello, WebAssembly!"))
