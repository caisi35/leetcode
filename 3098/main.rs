use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn sum_of_powers(nums: Vec<i32>, k: i32) -> i32 {
        const MOD : i64 = 1_000_000_007;
        const INF: i32 = 0x3f3f3f3f;
        let mut nums = nums.clone();
        let n = nums.len();
        let mut res = 0;
        let mut d: Vec<Vec<HashMap<i32, i64>>> = vec![vec![HashMap::new(); k as usize + 1]; n];
        nums.sort();
        for i in 0..n {
            d[i][1].insert(INF, 1);
            for j in 0..i {
                let diff = (nums[i] - nums[j]).abs();
                for p in 2..=k {
                    let mut updates: Vec<(i32, i64)> = Vec::new();
                    for (&v, &cnt) in &d[j][p as usize - 1] {
                        let key = diff.min(v);
                        updates.push((key, cnt));
                    }
                    for &(key, cnt) in &updates {
                        let entry = d[i][p as usize].entry(key).or_insert(0);
                        *entry = (*entry + cnt) % MOD;
                    }
                }
            }
            for (&v, &cnt) in &d[i][k as usize] {
                res = (res + v as i64 * cnt % MOD) % MOD;
            }
        }
        res as i32
    }
}

fn main() {
    println!("{}", Solution::sum_of_powers(vec![1,2,3,4], 3));
}