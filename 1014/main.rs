struct Solution {}

impl Solution {
    pub fn max_score_sightseeing_pair(values: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut mx = values[0] + 0;
        for j in 1..values.len() {
            ans = std::cmp::max(ans, mx + values[j as usize] - j as i32);
            mx = std::cmp::max(mx, values[j as usize] + j as i32);
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::max_score_sightseeing_pair(vec![1,2,3,4]));
}