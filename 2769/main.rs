
struct Solution{}

impl Solution {
    pub fn the_maximum_achievable_x(num: i32, t: i32) -> i32 {
        num + t * 2
    }
}

fn main() {
    println!("{}", Solution::the_maximum_achievable_x(4, 1))
}