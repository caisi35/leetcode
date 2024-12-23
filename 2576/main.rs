struct Solution {}

impl Solution {
    pub fn max_num_of_marked_indices(mut nums: Vec<i32>) -> i32 {
        nums.sort();
        let mut ans = 0;
        let n = nums.len();
        let m = n / 2;
        let mut i = 0;
        let mut j = m;
        while i < m && j < n {
            while j < n && 2 * nums[i] > nums[j] {
                j += 1;
            }
            if j < n {
                ans += 2;
                j += 1;
            }
            i += 1;
        }
        return ans;
    }
}

fn main() {
    println!("{}", Solution::max_num_of_marke_indices(vec![1,2,3,4]));
}