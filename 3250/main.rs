struct Solution{}

impl Solution {
    pub fn count_of_pairs(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let m = *nums.iter().max().unwrap();
        let mod_val = 1000000007;
        let mut dp = vec![vec![0; (m+1) as usize];n];
        for a in 0..=nums[0] {
            dp[0][a as usize] = 1;
        }
        for i in 1..n {
            let d = std::cmp::max(0, nums[i] - nums[i-1]);
            for j in d..=nums[i] {
                if j == 0 {
                    dp[i][j as usize] = dp[i-1][(j-d) as usize];
                } else {
                    dp[i][j as usize] = (dp[i][(j-1) as usize] + dp[i-1][(j-d) as usize]) % mod_val;
                }
            }
        }
        dp[n-1].iter().fold(0, |acc, &x| (acc+x) % mod_val)
    }
}

fn main() {
    println!("{}", Solution::count_of_pairs(vec![5,5,5,5]));
}