struct Solution{}

impl Solution {
    pub fn results_array(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let n = nums.len();
        let k = k as usize;
        let mut res = vec![-1;n-k+1];
        for i in 0..=n-k {
            let mut valid = true;
            for j in i+1..i+k {
                if nums[j] - nums[j-1] != 1 {
                    valid = false;
                    break;
                }
            }
            if valid {
                res[i] = nums[i+k-1]
            }
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::results_array(vec![3,2,3,2,3,2], 2));
}