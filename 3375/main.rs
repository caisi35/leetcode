use std::collections::HashSet;

struct Solution {}

impl Solution {
    pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
        let mut st = HashSet::new();
        for x in nums {
            if x < k {
                return -1
            } else if x > k {
                st.insert(x);
            }
        }
        st.len() as i32
    }
}

fn main() {
    println!("{}", Solution::min_operations(vec![2,1,2], 2));
}