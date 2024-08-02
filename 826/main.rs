use std::cmp::max;

struct Solution{}

impl Solution {
    pub fn max_profit_assignment(difficulty: Vec<i32>, profit: Vec<i32>, worker: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut dp: Vec<(i32, i32)> = difficulty.into_iter().zip(profit).collect();
        dp.sort_by_key(|&x| x.0);
        for w in worker.iter() {
            let mut m = 0;
            for &d in dp.iter() {
                if *w >= d.0 {
                    m = max(m, d.1)
                }
            }
            ans += m
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::max_profit_assignment(vec![85, 47, 57], vec![24, 66, 99], vec![40, 25, 25]))
}