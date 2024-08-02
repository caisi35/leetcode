use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn  relocate_marbles(nums: Vec<i32>, move_from: Vec<i32>, move_to: Vec<i32>) -> Vec<i32> {
        let mut m: HashMap<i32, i32> = HashMap::new();
        let mut ans: Vec<i32> = vec![];
        for n in nums.iter() {
            m.insert(*n, 1);
        }

        for i in 0..move_from.len() {
            if m.get(&(move_from[i] as i32)).is_some() {
                m.insert(move_from[i], 0);
                m.insert(move_to[i], 1);
            }
        }

        for (k, v) in m {
            if v >= 1 {
                ans.push(k)
            }
        }
        ans.sort();
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::relocate_marbles(vec![1,6,7,8], vec![1,7,2], vec![2,9,5]))
}