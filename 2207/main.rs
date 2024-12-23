struct Solution {}

impl Solution {
    pub fn maximum_subsequence_count(text: String, pattern: String) -> i64 {
        let mut res = 0;
        let mut cnt1 = 0;
        let mut cnt2 = 0;
        for c in text.chars() {
            if c == pattern.chars().nth(1).unwrap() {
                res += cnt1;
                cnt2 += 1;
            }
            if c == pattern.chars().nth(0).unwrap() {
                cnt1 += 1;
            }
        }
        res + std::cmp::max(cnt1, cnt2)
    }
}

fn main() {
    println!("{}", Solution::maximum_subsequence_count("abcdefj".to_string(), "ac".to_string()));
}