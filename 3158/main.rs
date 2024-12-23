use std::collections::HashSet;

struct Solution {}

impl Solution {
    pub fn duplicate_numbers_xor(nums: Vec<i32>) -> i32 {
        let mut cnt = HashSet::new();
        let mut ans = 0;
        for i in nums {
            if cnt.contains(&i) {
                ans ^= i;
            } else {
                cnt.insert(i);
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::duplicate_numbers_xor(vec![1,2,2,1]));
}