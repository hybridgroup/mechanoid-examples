#![no_std]

#[panic_handler]
fn handle_panic(_: &core::panic::PanicInfo) -> ! {
    core::arch::wasm32::unreachable()
}

#[no_mangle]
pub extern fn ping() {
    unsafe {pong()};
}

#[link(wasm_import_module = "hosted")]
extern {
    fn pong();
}
