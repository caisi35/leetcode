use std::cmp::{min, max};

struct Solution {}

impl Solution {
    pub fn min_moves_to_capture_the_queen(a: i32, b: i32, c: i32, d:i32, e: i32, f: i32) -> i32 {
        if a == e && (c != a || d <= min(b, f) || d >= max(b, f)) {
            return 1;
        }
        if b == f && (d != b || c <= min(a , e) || c >= max(a, e)) {
            return 1;
        }
        if (c - e).abs() == (d - f).abs() && ((c - e) * (b - f) != (a - e) * (d - f)
        || a < min(c, e) || a > max(c, e)) {
            return 1;
        }
        2
    }
}

fn main() {
    println!("{}", Solution::min_moves_to_capture_the_queen(1,1,8,8,2,3))
}