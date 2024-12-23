struct Solution {}

impl Solution {
    pub fn count_ways(mut nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut res = 0; 
        nums.sort();
        for k in 0..n+1 {
            if k > 0 && nums[k-1] >= k as i32 {
                continue
            }
            if k < n && nums[k] <= k as i32 {
                continue
            }
            res += 1;
        }
        return res
    }
}

fn main() {
    println!("{}", Solution::count_ways(vec![1, 1])) // 2
}