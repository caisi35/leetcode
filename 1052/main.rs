use std::cmp::max;

struct Solution {}

impl Solution {
    pub fn max_satisfied(customers: Vec<i32>, grumpy: Vec<i32>, minutes: i32) -> i32 {
        let mut total = 0;
        let n = customers.len() as usize;
        let minutes = minutes as usize;
        for i in 0..n {
            if grumpy[i] == 0 {
                total += customers[i];
            }
        }
        
        let mut increase = 0;
        for i in 0..minutes {
            let ii = i as usize;
            increase += customers[ii] * grumpy[ii];
        }

        let mut max_increase = increase;
        for i in minutes..n {
            let ii = i as usize;
            increase = increase - customers[ii-minutes] * grumpy[ii-minutes] + customers[ii] * grumpy[ii];
            max_increase = max(max_increase, increase);
        }
        total + max_increase
    }
}

fn main() {
    println!("{}", Solution::max_satisfied(vec![1, 0, 1, 2, 1, 1, 7, 5], vec![0,1,0,1,0,1,0,1], 3))
}