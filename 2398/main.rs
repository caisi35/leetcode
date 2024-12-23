use std::collections::VecDeque;

struct Solution {}

impl Solution {
    pub fn maximum_robots(charge_times: Vec<i32>, running_costs:Vec<i32>, budget: i64) -> i32 {

        let mut res = 0;
        let n = charge_times.len();
        let mut running_const_sum: i64 = 0;
        let mut q: VecDeque<usize> = VecDeque::new();
        let mut j = 0;
        for i in 0..n {
            running_const_sum += running_costs[i] as i64;
            while !q.is_empty() && charge_times[*q.back().unwrap()] <= charge_times[i] {
                q.pop_back();
            }
            q.push_back(i);
            while j <= i && (i - j + 1) as i64 * running_const_sum + charge_times[*q.front().unwrap()] as i64 > budget {
                if !q.is_empty() &&  *q.front().unwrap() == j {
                    q.pop_front();
                }
                running_const_sum -= running_costs[j] as i64;
                j += 1;
            }
            res = res.max(i - j + 1);
        }
        return res as i32
    }
}

fn main() {
    println!("{}", Solution::maximum_robots(vec![1,2,3,4,5], vec![1,2,3,4,5], 25))
}