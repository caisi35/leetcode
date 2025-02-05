struct Solution {}

impl Solution {
    pub fn sort_array_by_parity_ii(nums: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![0; nums.len()];
        let mut i = 0;
        for &n in &nums {
            if n % 2 == 0 {
                ans[i] = n;
                i += 2;
            }
        }
        i = 1;
        for &n in &nums {
            if n % 2 == 1 {
                ans[i] = n;
                i += 2;
            }
        }
        return ans;
    }
}

fn main() {
    println!("{:#?}", Solution::sort_array_by_parity_ii(vec![4,2,5,7]));
}