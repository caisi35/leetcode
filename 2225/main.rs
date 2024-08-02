use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn find_winners(matches: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut m = HashMap::new();
        for matche in matches {
            let (w, l) = (matche[0], matche[1]);
            m.entry(w).or_insert(0);
            *m.entry(l).or_insert(0) += 1;
        }

        let mut ans = vec![Vec::new(), Vec::new()];
        for (key, value) in m {
            if value < 2 {
                ans[value as usize].push(key);
            }
        }
        ans[0].sort_unstable();
        ans[1].sort_unstable();
        ans
    }
}

fn main() {
    println!(
        "{:#?}",
        Solution::find_winners(vec![
            vec![1, 3],
            vec![2, 3],
            vec![3, 6],
            vec![5, 6],
            vec![5, 7],
            vec![4, 5],
            vec![4, 8],
            vec![4, 9],
            vec![10, 4],
            vec![10, 9]
        ])
    )
}
