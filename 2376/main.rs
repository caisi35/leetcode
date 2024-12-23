use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn count_special_numbers(n: i32) -> i32 {
        let n_str = n.to_string();
        let mut memo: HashMap<i32, i32> = HashMap::new();
        fn dp(mask: i32, prefix_smaller: bool, n_str: &str, memo: &mut HashMap<i32, i32>) -> i32 {
            if mask.count_ones() as usize == n_str.len() {
                return 1;
            }
            let key = mask * 2 + if prefix_smaller {1} else {0};
            if !memo.contains_key(&key) {
                let mut res = 0;
                let lower_bound = if mask == 0 {1} else {0};
                let upper_bound = if prefix_smaller {
                    9
                } else {
                    (n_str.chars().nth(mask.count_ones() as usize).unwrap() as i32) - ('0' as i32)
                };
                for i in lower_bound..=upper_bound {
                    if (mask >> i) & 1 == 0 {
                        res += dp(mask | (1<<i), prefix_smaller || i < upper_bound, n_str, memo);
                    }
                }
                memo.insert(key, res);
            }
            *memo.get(&key).unwrap()
        }
        let mut res = 0;
        let mut prod = 9;
        for i in 0..(n_str.len()-1) {
            res += prod;
            prod *= 9-i as i32;
        }
        res += dp(0, false, &n_str, &mut memo);
        res
    }
}

fn main() {
    println!("{}", Solution::count_special_numbers(135));
}