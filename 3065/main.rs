struct Solution{}

impl Solution {
    pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
        let mut ans = 0;
        for i in 0..nums.len() {
            if nums[i] < k {
                ans += 1;
            }
        }
        ans
    }
}

fn main() {
    println("{:#?}", Solution::min_operations(vec![1,1,2,4,9], 1));
}