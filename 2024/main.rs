use std::cmp::max;

struct Solution {}

impl Solution {
    pub fn max_consecutive_answers(answer_key: String, k: i32) -> i32 {
        max(Self::max_consecutive_char(answer_key.clone(), k, 'T'), Self::max_consecutive_char(answer_key, k, 'F'))
    }

    fn max_consecutive_char(answer_key: String, k: i32, ch: char) -> i32 {
        let mut ans = 0;
        let mut left = 0;
        let mut sum = 0;
        let ans_key = answer_key.as_bytes();
        for (right, _) in ans_key.iter().enumerate() {
            if ans_key[right as usize] != ch as u8 {
                sum += 1;
            }
            while sum > k {
                if ans_key[left as usize] != ch as u8 {
                    sum -= 1;
                }
                left += 1;
            }
            ans = max(ans, (right as usize)-left+1)
        }
        ans as i32
    }
}

fn main() {
    println!("{}", Solution::max_consecutive_answers("TTFF".to_string(), 2)); // 4
    println!("{}", Solution::max_consecutive_answers("TFFT".to_string(), 1)); // 3
}
