struct Solution{}

impl Solution {
    pub fn count_subarrays(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut ans = 0;
        for i in 1..n-1 {
            if nums[i] == (nums[i-1] + nums[i+1]) * 2 {
                ans += 1;
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::count_subarrays(vec![-1,-4,-1,4])) // 1
}