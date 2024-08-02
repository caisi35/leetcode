struct Solution {}

impl Solution {
    pub fn c(nums: Vec<i32>, target: i32) -> i32 {
        let t = target as usize;
        let mut dp = vec![0; t+1];
        dp[0] = 1;
        for i in 1..=t {
            for &num in nums.iter() {
                let x = num as usize;
                // println!("{} {} ", x, i);
                if x <= i {
                    dp[i] += dp[i-x];
                }
            }
        }
        // println!("{}", dp);
        dp[t]
    }
}

fn main() {
    println!("{}", Solution::c(vec![1, 2, 3], 4))
}