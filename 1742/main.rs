use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn count_balls(low_limit: i32, high_limit: i32) -> i32 {
        let mut c = HashMap::<i32, i32>::new();
        let mut ans = 0;
        for i in low_limit..=high_limit {
            let mut sum = 0;
            let mut x = i;
            while x > 0 {
                sum += x % 10;
                x /= 10;
            }
            if let Some(_s) = c.get(&sum) {
                c.insert(sum, _s + 1);
            } else {
                c.insert(sum, 1);
            }
            ans = ans.max(c[&sum]);
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::count_balls(1, 10));
}