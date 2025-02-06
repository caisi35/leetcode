struct Solution{}

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> bool {
        let n = nums.len();
        if n == 0 {
            return false;
        }
        if n == 1 {
            return nums[0] == target;
        }
        let (mut l, mut r) = (0, n as i32 - 1);
        while l <= r {
            let mid = l + (r - l) / 2;
            if nums[mid as usize] == target {
                return true
            }
            if nums[l as usize] == nums[mid as usize] && nums[mid as usize] == nums[r as usize] {
                l += 1;
                r -= 1;
            } else if nums[l as usize] <= nums[mid as usize] {
                if nums[l as usize] <= target && target < nums[mid as usize] {
                    r = mid - 1;
                } else {
                    l = mid + 1;
                }
            } else {
                if nums[mid as usize] < target && target <= nums[n-1] {
                    l = mid + 1;
                } else {
                    r = mid - 1;
                }
            }
        }
        false
    }
}

fn main() {
    println!("{:#?}", Solution::search(vec![2,5,6,0,0,1,2], 0));
}