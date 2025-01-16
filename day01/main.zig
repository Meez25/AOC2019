const std = @import("std");
const time = std.time;
const data = @embedFile("input.txt");

pub fn main() !void {
    const start = try std.time.Instant.now();

    // Preallocate an ArrayList for better memory management
    var total: i32 = 0;
    var i: usize = 0;

    // Direct iteration over bytes instead of splitting
    while (i < data.len) {
        // Skip empty lines and whitespace
        if (data[i] == '\n') {
            i += 1;
            continue;
        }

        // Parse number directly
        var num: i32 = 0;
        while (i < data.len and data[i] != '\n') {
            num = num * 10 + (data[i] - '0');
            i += 1;
        }

        // Calculate fuel
        total += @divTrunc(num, 3) - 2;
    }

    const end = try std.time.Instant.now();
    const elapsed_ns = end.since(start);
    const elapsed_us = @as(f64, @floatFromInt(elapsed_ns)) / 1_000.0;
    std.debug.print("Part one: {d}, in {d:.3}us\n", .{ total, elapsed_us });
}
