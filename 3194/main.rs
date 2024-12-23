struct Solution {}

impl Solution {
    pub fn minimum_average(nums: Vec<i32>) -> f64 {
        let mut nums = nums;
        nums.sort();
        let n = nums.len();
        let mut res = f64::INFINITY;
        for i in 0..n/2 {
            let avg = (nums[i] as f64 + nums[n-i-1] as f64) / 2.0;
            res = res.min(avg);
        }
        res
    }
}

fn main() {
    println!("{}", Solution::minimum_average(vec![1,2,3,4,5,6]));
}