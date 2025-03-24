struct Solution{}

impl Solution {
    pub fn count_prefixes(words: Vec<String>, s: String) -> i32 {
        let mut ans = 0;
        for word in words {
            if s.starts_with(&word) {
                ans += 1;
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::count_prefixes(vec!["a".to_string(), "a".to_string()], "aa".to_string()));
}