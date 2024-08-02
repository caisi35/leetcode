
struct Solution {}

impl Solution {
    pub fn most_competitive(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut res= Vec::new();
        let n = nums.len();
        for i in 0..n {
            while res.len() > 0 && (n - i + res.len()) as i32 > k && *res.last().unwrap() > nums[i] {
                res.pop();
            }
            res.push(nums[i]);
        }
        res.truncate(k as usize);
        res
    }
}

fn main() {
    println!("{:#?}", Solution::most_competitive(vec![3,5,2,6], 2))
}