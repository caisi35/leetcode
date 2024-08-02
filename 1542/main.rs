use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn longest_awesome(s: String) -> i32 {
        let mut prefix: HashMap<i32, i32> = HashMap::new();
        prefix.insert(0, -1);
        let mut ans = 0;
        let mut sequence = 0;
        for (j, ch) in s.chars().enumerate() {
            let digit = ch.to_digit(10).unwrap() as i32;
            sequence ^= 1 << digit;
            if let Some(&prev_index) = prefix.get(&sequence) {
                ans = ans.max(j as i32 - prev_index);
            } else {
                prefix.insert(sequence, j as i32);
            }
            for k in 0..10 {
                if let Some(&prev_index) = prefix.get(&(sequence ^ (1 << k))) {
                    ans = ans.max(j as i32 - prev_index);
                }
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::longest_awesome("3242415".to_string()))
}