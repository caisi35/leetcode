struct Solution {}

impl Solution {
    pub fn results_array(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let n = nums.len();
        let mut cnt = 0;
        let k = k as usize;
        let mut ans = vec![-1; n-k+1];
        for i in 0..n {
            if i == 0 || nums[i] - nums[i-1] != 1 {
                cnt = 1
            } else {
                cnt += 1
            }
            if cnt >= k {
                ans[i-k+1] = nums[i]
            }
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::results_array(vec![3,2,3,2,3,2], 2));
}