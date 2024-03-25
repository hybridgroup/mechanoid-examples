
extern "hosted" fn pong() void;

pub export fn ping() void {
    pong();
}
