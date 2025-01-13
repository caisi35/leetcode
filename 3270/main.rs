use std::cmp::min;

struct Solution {}

impl Solution {
    pub fn generate_key(mut num1: i32, mut num2: i32, mut num3: i32) -> i32 {
        let mut key = 0;
        let mut p = 1;
        while num1 > 0 && num2 > 0 && num3 > 0 {
            key += min(min(num1 % 10, num2 % 10), num3 % 10) * p;
            num1 /= 10;
            num2 /= 10;
            num3 /= 10;
            p *= 10;
        }
        key
    }
}

fn main() {
    println!("{:#?}", Solution::generate_key(1, 10, 1000));
}