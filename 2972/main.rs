struct Solution{}

impl Solution {
    pub fn incremovable_subarray_count(nums: Vec<i32>) -> i64 {
        let n = nums.len();
        let mut ans: i64 = 0;
        let mut l: i32 = 0;

        while (l as usize) < n - 1 {
            if nums[l as usize] >= nums[l as usize+1] {
                break;
            }
            l += 1;
        }
        if l as usize == n-1 {
            return (n * (n + 1) / 2) as i64;
        }
        ans += (l+2) as i64;
        let mut r = n-1;
        while r > 0 {
            if r < n-1 && nums[r] >= nums[r+1] {
                break;
            }
            
            while l as i32 >= 0 && nums[l as usize] >= nums[r] {
                l -= 1;
            }
            ans += (l+2) as i64;
            r -= 1;
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::incremovable_subarray_count(vec![1,2,3,4]));
    println!("{}", Solution::incremovable_subarray_count(vec![8,7,6,6]));

}