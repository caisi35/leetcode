struct Solution {}

impl Solution {
    pub fn number_of_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i32 {
        let mut ans = 0;
        for i in nums1.iter() {
            for j in nums2.iter() {
                if i % (j * k) == 0 {
                    ans += 1;
                }
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::number_of_pairs(vec![1,3,4], vec![1,3,4], 1))
}