use std::cmp::max;

struct Solution {}

impl Solution {
    pub fn temperature_trend(temperature_a: Vec<i32>, temperature_b: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut f = 0;
        for i in 1..temperature_a.len() {
            if Self::c(temperature_a[i]-temperature_a[i-1], temperature_b[i]-temperature_b[i-1]) {
                f += 1;
            } else {
                f = 0;
            }
            ans = max(ans, f);
        }
        ans
    }

    fn c(a: i32, b: i32) -> bool {
        if a == b {
            return a == b 
        } else if a > 0 && b > 0 {
            return true
        } else if a < 0 && b < 0 {
            return true
        }
        return false
    }
}




fn main() {
    println!("{}", Solution::temperature_trend(vec![21,18,18,18,31], vec![34,32,16,16,17]))
}