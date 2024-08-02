struct Solution{}

impl Solution {
    pub fn minimum_steps(s: String) -> i64 {
        let mut ans: i64 = 0;
        let mut sum = 0;
        for i in s.chars() {
            if i == '1' {
                sum += 1;
            } else {
                ans += sum as i64;
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::minimum_steps("101".to_string()))
}