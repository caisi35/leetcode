use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn number_of_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i64 {
        let mut ans: i64 = 0;
        let mut c1: HashMap<i32, i64>= HashMap::new();
        let mut c2: HashMap<i32, i64>= HashMap::new();
        let mut m = 0 ;
        for i in nums1.iter() {
            *c1.entry(*i).or_insert(0) += 1;
            if *i > m {
                m = *i as i32;
            }
        }
        for j in nums2.iter() {
            *c2.entry(*j).or_insert(0) += 1;
        }
    
        for (&a, &v) in &c2 {
            for b in (a * k..=m).step_by((a * k) as usize) {
                if let Some(&va) = c1.get(&b) {
                    ans += v * va;
                }
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::number_of_pairs(vec![1,2,3,4,5], vec![2,3,4,5,6], 2));
}