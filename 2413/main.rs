struct Solution {}

impl Solution {
    pub fn smallest_even_multiple(n: i32) -> i32 {
        if n % 2 == 0 {
            return n
        } else {
            return 2 * n
        }
    }
}

fn main() {
    println!("{}", Solution::smallest_even_multiple(5))
}