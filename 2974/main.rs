struct Solution {}

impl Solution {
    pub fn number_game(mut nums: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![0; nums.len()];
        nums.sort();
        let mut i = 1;
        while i < nums.len() {
            ans[i] = nums[i-1];
            ans[i-1] = nums[i];
            i += 2;
        }

        return ans;
    }
}

fn main() {
    println!("{:#?}", Solution::number_game(vec![5,4,2,3]))
}