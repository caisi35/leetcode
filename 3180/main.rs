struct Solution {}

impl Solution {
    pub fn max_total_reward(reward_values: Vec<i32>) ->i32 {
        let mut reward_values = reward_values.clone();
        reward_values.sort();
        let m = reward_values[reward_values.len()-1] as usize;
        let mut dp = vec![0; 2 * m];
        dp[0] = 1;
        for &x in &reward_values {
            let x = x as usize;
            for k in (x..2 * x).rev() {
                if dp[k-x] == 1 {
                    dp[k] = 1
                }
            }
        }
        let mut res = 0;
        for (i, &value) in dp.iter().enumerate() {
            if value == 1 {
                res = i;
            }
        }
        res as i32
    }
}

fn main() {
    println!("{}", Solution::max_total_reward(vec![1,1,3,3]));
}