struct Solution {}

impl Solution {
    pub fn get_final_state(mut nums: Vec<i32>, k: i32, multiplier: i32) -> Vec<i32> {
        for _ in 0..k {
            let mut m = 0;
            for j in 1..nums.len() {
                if nums[j] < nums[m] {
                    m = j;
                }
            }
            nums[m] *= multiplier;
        }
        nums
    }
}

fn main() {
    println!("{:#?}", Solution::get_final_state(vec![1,2], 3, 4));
}