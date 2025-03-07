use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn beautiful_subsets(nums: Vec<i32>, k: i32) -> i32 {
        let mut ans = 0;
        let mut cnt = HashMap::new();
        fn dfs(i: usize, nums: &Vec<i32>, k: i32, ans: &mut i32, cnt: &mut HashMap<i32, i32>) {
            if i == nums.len() {
                *ans += 1;
                return;
            }
            dfs(i+1, nums, k, ans, cnt);
            if cnt.get(&(nums[i]-k)).unwrap_or(&0) == &0 && cnt.get(&(nums[i]+k)).unwrap_or(&0) == &0 {
                *cnt.entry(nums[i]).or_insert(0) += 1;
                dfs(i+1, nums, k, ans, cnt);
                *cnt.entry(nums[i]).or_insert(0) -= 1;
            }
        }
        dfs(0, &nums, k, &mut ans, &mut cnt);
        ans - 1
    }
}

fn main() {
    println!("{:#?}", Solution::count_beautiful_subsets(vec![1,2,3,4,5], 2));
}