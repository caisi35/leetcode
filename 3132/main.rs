struct Solution{}

impl Solution {
    pub fn minimum_added_integer(mut nums1: Vec<i32>, mut nums2: Vec<i32>) -> i32 {
        nums1.sort();
        nums2.sort();
        for i in (0..=2).rev() {
            let (mut left, mut right) = (i as usize + 1, 1);
            while left < nums1.len() && right < nums2.len() {
                if nums1[left] - nums2[right] == nums1[i as usize] - nums2[0] {
                    right += 1;
                }
                left += 1;
            }
            if right == nums2.len() {
                return nums2[0] - nums1[i as usize];
            }
        }
        0
    }
}

fn main() {
    println!("{}", Solution::minimum_added_integer(vec![4,20,16,12,8], vec![14,18,10]))
}