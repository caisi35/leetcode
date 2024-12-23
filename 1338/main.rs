use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn min_set_size(arr: Vec<i32>) -> i32 {
        let mut freq = HashMap::new();
        for &num in &arr {
            *freq.entry(num).or_insert(0) += 1;
        }
        let mut occ: Vec<i32> = freq.values().cloned().collect();
        occ.sort_by(|a, b| b.cmp(a));
        let mut cnt = 0;
        let mut ans = 0;
        for c in occ {
            cnt += c;
            ans += 1;
            if cnt * 2 >= arr.len() as i32 {
                break;
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::min_set_size(vec![7,7,7,7,7,7]));
}