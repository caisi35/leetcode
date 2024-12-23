struct Solution {}

impl Solution {
    pub fn clear_digits(s: String) -> String {
        let mut ans: Vec<char> = Vec::new();
        for ch in s.chars() {
            if ch.is_digit(10) {
                ans.pop();
            } else {
                ans.push(ch);
            }
        }
        ans.into_iter().collect()
    }
}

fn main() {
    println!("{}", Solution::clear_digits("d9".to_string()));
}