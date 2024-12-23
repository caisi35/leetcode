struct Solution{}

fn main() {
	println!("{}", Solution::count_quadruplets(vec![1,3,2,4,5]))
}

impl Solution {
    pub fn count_quadruplets(nums: Vec<i32>) -> i64 {
        let n = nums.len();
        let mut pre = vec![0; n + 1];
        let mut ans: i64 = 0;
        for j in 0..(n as usize) {
            let mut suf = 0;
            for k in (j+1..n).rev() {
                if nums[j] > nums[k] {
                    ans += pre[nums[k as usize] as usize] as i64 * suf as i64
                } else {
                    suf += 1
                }
            }
            for x in (nums[j] as usize + 1)..(n as usize) {
                pre[x as usize] += 1;
            }
        }
        return ans;
    }
}