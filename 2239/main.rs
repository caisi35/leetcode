struct Solution {}

impl Solution {
    pub fn find_closest_number(nums: Vec<i32>) -> i32 {
        let mut res = nums[0];
        let mut dis = nums[0].abs();
        for &num in nums.iter() {
            if num.abs() < dis {
                dis = num.abs();
                res = num;
            } else if num.abs() == dis {
                res = res.max(num);
            }
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::find_closest_number(vec![-1, -1]));
}