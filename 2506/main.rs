use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn similar_pairs(words: Vec<String>) -> i32 {
        let mut res = 0;
        let mut cnt = HashMap::new();
        for word in words {
            let mut state = 0;
            for c in word.chars() {
                state |= 1 << (c as u8 - b'a');
            }
            res += cnt.get(&state).unwrap_or(&0);
            *cnt.entry(state).or_insert(0) += 1;
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::similar_pairs(vec!["ab".to_string(), "ba".to_string()]));
}