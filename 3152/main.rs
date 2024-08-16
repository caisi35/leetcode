struct Solution{}

impl Solution {
    pub fn is_array_special(nums: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<bool> {
        let n = nums.len();
        let mut dp: Vec<i32> = vec![0; n];
        for i in 0..n {
            dp[i as usize] = 1;
        }
        for i in 1..n {
            if (nums[i as usize] ^ nums[i as usize - 1]) & 1 == 1 {
                dp[i as usize] = dp[i as usize - 1] + 1;
            }
        }
        let mut res: Vec<bool> = vec![false; queries.len()];
        for (i, q) in queries.iter().enumerate() {
            let (x, y) = (q[0], q[1]);
            res[i as usize] = dp[y as usize] >= y - x + 1;
        }
        res
    }
}

fn main() {
    println!("{:?}", Solution::is_array_special(vec![3,3,2], vec![vec![1,2]]))
}