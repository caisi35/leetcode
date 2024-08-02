struct Solution {}

impl Solution {
    pub fn incremovable_subarray_count(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut res = 0;
        for i in 0..n {
            for j in i..n {
                if Self::is_increasing(&nums, i, j) {
                    res += 1;
                }
            }
        }
        res
    }

    fn is_increasing(nums: &Vec<i32>, l: usize, r: usize) -> bool {
        for i in 1..nums.len() {
            if i >= l && i <= r+1 {
                continue;
            }
            if nums[i] <= nums[i-1] {
                return false;
            }
        }
        if l >= 1 && r+1 < nums.len() && nums[r+1] <= nums[l-1] {
            return false;
        }
        true
    }
}

fn main() {
    println!("{}", Solution::incremovable_subarray_count(vec![1,2,3,4]))
}