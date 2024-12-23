struct Solution{}

impl Solution {
    pub fn longest_continuous_substring(s: String) -> i32 {
        let mut cur = 1;
        let mut ans = 1;
        for i in 1..s.len() {
            if s.as_bytes()[i] == s.as_bytes()[i-1] + 1 {
                cur+=1;
            } else {
                cur = 1;
            }
            ans = std::cmp::max(cur, ans);
        }
        return ans
    }
}

fn main() {
    println!("{}", Solution::longest_continuous_substring("abdaabc".to_string()));
}