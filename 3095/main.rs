struct Solution{}

impl Solution {
    pub fn minimum_subarray_length(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        let mut res = n + 1;
        for i in 0..n {
            let mut v = 0;
            for j in i..n {
                v |= nums[j];
                if v >= k {
                    res = std::cmp::min(res, j-i+1);
                    break
                }
            }
        }
        if res == n+1 {
            return -1
        }
        res as i32
    }
}

fn main() {
    println!("{:#?}", Solution::minimum_subarray_length(vec![1,2,3], 2));
}