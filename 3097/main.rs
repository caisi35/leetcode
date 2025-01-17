struct Solution{}

impl Solution {
    pub fn minimum_subarray_length(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        let mut bits = [0;30];
        let mut res = i32::MAX;

        let calc = |bits: &[i32]| -> i32 {
            let mut ans = 0;
            for i in 0..30 {
                if bits[i] > 0 {
                    ans |= 1 << i;
                }
            }
            ans
        };
        let mut left = 0;
        for right in 0..n {
            for i in 0..30 {
                bits[i] += (nums[right] >> i) & 1;
            }
            while left <= right && calc(&bits) >= k {
                res = res.min((right-left+1) as i32);
                for i in 0..30 {
                    bits[i] -= (nums[left] >> i) & 1;
                }
                left += 1;
            }
        }
        if res == i32::MAX {
            -1
        } else {
            res
        }
    }
}

fn main() {
    println!("{:#?}", Solution::minimum_subarray_length(vec![1,2,3], 2));
}