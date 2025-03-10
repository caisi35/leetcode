struct Solution {}

impl Solution {
    pub fn divisor_substrings(num: i32, k: i32) -> i32 {
        let s = num.to_string();
        let n = s.len();
        let mut res = 0;
        for i in 0..= n - k as usize {
            let tmp: i32 = s[i..i+k as usize].parse().unwrap();
            if tmp != 0 && num % tmp == 0 {
                res += 1;
            }
        }
        res
    }
}

fn main() {
    println!("{}", Solution::divisor_substrings(240, 2));
}