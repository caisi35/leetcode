struct Solution {}

impl Solution {
    pub fn added_integer(mut nums1: Vec<i32>, mut nums2: Vec<i32>) -> i32 {
        nums1.sort();
        nums2.sort();
        return -(nums1[0] - nums2[0])
    }
}

fn main() {
    println!("{}", Solution::added_integer(vec![2,6,4], vec![9,7,5]))
}