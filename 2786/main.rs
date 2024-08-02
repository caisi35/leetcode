struct Solution{}

impl Solution {
    pub fn max_score(nums: Vec<i32>, x: i32) -> i64 {
        let mut res = nums[0] as i64;
        let mut dp: Vec<i64> = vec![i32::MIN as i64, i32::MIN as i64];
        dp[(nums[0] % 2) as usize] = nums[0] as i64;
        for i in 1..nums.len() {
            let par = (nums[i] % 2) as usize;
            let cur = (dp[par] + nums[i] as i64).max(dp[1-par] - x as i64 + nums[i] as i64);
            res = res.max(cur);
            dp[par] = dp[par].max(cur);
        }
        return res;
    }
}

fn main() {
    println!("{}", Solution::max_score(vec![2,3,6,1,9,2], 5))
}