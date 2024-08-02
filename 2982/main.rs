use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn maximum_length(s: String) -> i32 {
        let n = s.len();
        let mut cnt = HashMap::new();
        let s_bytes = s.as_bytes();

        let mut i = 0;
        while i < n {
            let mut j = i;
            while j < n && s_bytes[j] == s_bytes[i] {
                j += 1;
            }
            cnt.entry(s_bytes[i])
                .or_insert_with(Vec::new)
                .push((j-i) as i32);
            i = j;
        }
        let mut res = -1;
        for vec in cnt.values() {
            let mut lo = 1;
            let mut hi = (n as i32) - 2;
            while lo <= hi {
                let mid = (lo + hi) / 2;
                let mut count = 0;
                for &x in vec {
                    if x >= mid {
                        count += x - mid + 1;
                    }
                }
                if count >= 3 {
                    res = res.max(mid);
                    lo = mid + 1;
                } else {
                    hi = mid - 1;
                }
            }
        }
        res
    }
}

fn main() {
    println!("{}", Solution::maximum_length("abcaba".to_string()))
}