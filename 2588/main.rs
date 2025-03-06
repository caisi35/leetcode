use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn beautiful_subarrays(nums: Vec<i32>) -> i64 {
        let mut cnt = HashMap::new();
        cnt.insert(0, 1);
        let mut mask = 0;
        let mut ans = 0;
        for x in nums {
            mask ^= x;
            ans += *cnt.get(&mask).unwrap_or(&0);
            *cnt.entry(mask).or_insert(0) += 1;
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::beautiful_subarrays(vec![1,2,3,4,5]));
}