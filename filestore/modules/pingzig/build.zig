const std = @import("std");

pub fn build(b: *std.Build) void {
    const lib = b.addSharedLibrary(.{
        .name = "pingzig",
        .root_source_file = .{ .path = "src/ping.zig" },
        .target = .{
            .cpu_arch = .wasm32,
            .os_tag = .freestanding,
        },
        .optimize = .ReleaseSmall,
        .link_libc = true,
    });
    lib.rdynamic = true;
    lib.stack_size = 4096;
    lib.initial_memory = 65536;
    lib.max_memory = 65536;

    b.installArtifact(lib);
}
