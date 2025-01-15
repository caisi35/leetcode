use std::collections::BinaryHeap;
use std::cmp::Reverse;

struct Solution{}

impl Solution {
    pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
        let mut res = 0;
        let mut pq: BinaryHeap<Reverse<i64>> = BinaryHeap::new();
        for &num in &nums {
            pq.push(Reverse(num as i64));
        }
        while let Some(Reverse(x)) = pq.pop() {
            if x >= k as i64 {
                break;
            }
            if let Some(Reverse(y)) = pq.pop() {
                pq.push(Reverse(x+x+y));
                res += 1;
            }
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::min_operations(vec![2,11,10,1,3], 10));
}