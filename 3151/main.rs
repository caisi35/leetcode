
struct Solution{}

impl Solution {
    pub fn is_array_special(nums: Vec<i32>) -> bool {
        if nums.len() == 1 {
            return true
        }

        for i in 0..nums.len()-1 {
            if (nums[i+1] + nums[i]) % 2 == 0 {
                return false
            }
        }
        true
    }
}

fn main() {
    println!("{}", Solution::is_array_special(vec![4,1,2,3]))
}