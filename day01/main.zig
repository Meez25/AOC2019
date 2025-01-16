const std = @import("std");
const time = std.time;
const data = @embedFile("input.txt");

pub fn main() !void {
    const start = try std.time.Instant.now();

    // Preallocate an ArrayList for better memory management
    var total: i32 = 0;
    var totalp2: i32 = 0;
    var i: usize = 0;

    while (i < data.len) {
        if (data[i] == '\n') {
            i += 1;
            continue;
        }

        var num: i32 = 0;
        while (i < data.len and data[i] != '\n') {
            num = num * 10 + (data[i] - '0');
            i += 1;
        }

        total += @divTrunc(num, 3) - 2;
        totalp2 += compute_fuel(num, 0);
    }

    const end = try std.time.Instant.now();
    const elapsed_ns = end.since(start);
    const elapsed_us = @as(f64, @floatFromInt(elapsed_ns)) / 1_000.0;
    std.debug.print("Part one: {d}, Part two : {d} in {d:.3}us\n", .{ total, totalp2, elapsed_us });
}

pub fn compute_fuel(fuel: i32, total: i32) i32 {
    if (fuel <= 0) return total;
    const fuelNeeded = @divTrunc(fuel, 3) - 2;
    if (fuelNeeded <= 0) return total;
    return compute_fuel(fuelNeeded, total + fuelNeeded);
}
