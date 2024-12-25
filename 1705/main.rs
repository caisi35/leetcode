use std::collections::BinaryHeap;
use std::cmp::Reverse;

struct Solution{}

impl Solution {
    pub fn eaten_apples(apples: Vec<i32>, days: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut pq = BinaryHeap::new();
        let n = apples.len();
        let mut i = 0;
        while i < n {
            while let Some(Reverse((rotten_day, _count))) = pq.peek() {
                if *rotten_day <= i as i32 {
                    pq.pop();
                } else {
                    break;
                }
            }
            let rotten_day = i as i32 + days[i];
            let _count = apples[i];
            if _count > 0 {
                pq.push(Reverse((rotten_day, _count)));
            }
            if let Some(Reverse((rotten_day, mut _count))) = pq.pop() {
                _count -= 1;
                if _count > 0 {
                    pq.push(Reverse((rotten_day, _count)));
                }
                ans += 1;
            }
            i += 1;
        }
        while let Some(Reverse((rotten_day, c))) = pq.pop() {
            if rotten_day <= i as i32 {
                continue
            }
            let num = std::cmp::min(rotten_day - i as i32, c);
            ans += num;
            i += num as usize;
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::eaten_apples(vec![1,2,3,5,2], vec![3,2,1,4,2]));
}