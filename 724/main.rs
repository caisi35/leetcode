struct Solution {}

impl Solution {
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        let mut t: i32 = 0;
        for i in 0..nums.len() {
            t += nums[i];
        }

        let mut s: i32 = 0;
        for (i, &v) in nums.iter().enumerate() {
            if 2 * s + v == t {
                return i as i32
            }
            s += v;
        }
        -1
    }
}

fn main() {
    println!("{}", Solution::pivot_index(vec![1,7,3,6,5,6]))
}