struct Solution {}

impl Solution {
    pub fn sum_digit_differences(nums: Vec<i32>) -> i64 {
        let mut nums = nums.clone();
        let mut res = 0i64;
        while nums[0] > 0 {
            let mut cnt: Vec<i32> = vec![0; 10];
            for i in 0..nums.len() {
                cnt[(nums[i as usize] % 10) as usize] += 1;
                nums[i as usize] /= 10;
            }
            for i in 0..10 {
                res += (nums.len() as i32 - cnt[i as usize]) as i64 * cnt[i as usize] as i64;
            }
        }
        return res / 2;
    }
}

fn main() {
    println!("{}", Solution::sum_digit_differences(vec![12,23,13]))  // 4
}