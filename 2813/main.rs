use std::collections::{VecDeque,HashSet} ;

struct Solution{}

impl Solution {
    pub fn find_maximum_elegance(mut items: Vec<Vec<i32>>, k: i32) -> i64 {
        items.sort_unstable_by_key(|item| -item[0]);
        let (mut ca, mut st) = (HashSet::new(), VecDeque::new());
        let ( mut res, mut profix) = (0 as i64, 0 as i64);
        for (i, item) in items.iter().enumerate() {
            if i < k as usize {
                profix += item[0] as i64;
                if !ca.contains(&item[1]) {
                    ca.insert(item[1]);
                } else {
                    st.push_back(item[0]);
                }
            } else if !ca.contains(&item[1]) && !st.is_empty() {
                profix += (item[0] - st.back().unwrap()) as i64;
                st.pop_back();
                ca.insert(item[1]);
            }
            res = res.max(profix + (ca.len() * ca.len()) as i64);
        }
        res as i64
    }
}

fn main() {
    println!("{}", Solution::find_maximum_elegance(vec![vec![3,2],vec![5,1],vec![10,1]], 2))
}