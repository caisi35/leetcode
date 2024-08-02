struct Solution {}

impl Solution {
    pub fn max_operations(nums: Vec<i32>) -> i32 {
        let mut ans = 1;
        let s = nums[0] + nums[1];
        let mut c = 2;
        while c < nums.len() - 1 {
            if nums[c] + nums[c + 1] == s {
                ans += 1;
            } else {
                break;
            }
            c += 2;
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::max_operations(vec![3, 2, 1, 4, 5]));
    println!("{}", Solution::max_operations(vec![3, 2, 6, 1, 4]));
    println!(
        "{}",
        Solution::max_operations(vec![2, 2, 3, 2, 4, 2, 3, 3, 1, 3])
    );
    println!("{}", Solution::max_operations(vec![3, 2, 1, 4, 1]));

}
