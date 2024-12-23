struct Solution {}

impl Solution {
    pub fn smallest_range_ii(mut nums: Vec<i32>, k: i32) -> i32 {
        nums.sort();
        let mut r = nums[nums.len()-1] - nums[0];
        for i in 0..nums.len()-1 {
            let (a , b) = (nums[i], nums[i+1]);
            r = std::cmp::min(r, std::cmp::max(nums[nums.len()-1]-k, a+k)-std::cmp::min(nums[0]+k, b-k))
        }
        r
    }
}

fn main() {
    println!("{}", Solution::smallest_range_ii(vec![1,3,6], 3)); // 3
}