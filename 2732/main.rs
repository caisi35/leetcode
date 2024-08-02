use std::collections::HashMap;
struct Solution{}

impl Solution {
    pub fn good_subsetof_binary_matrix(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ans: Vec<i32> = vec![];
        let mut mp:HashMap<i32, i32> = HashMap::new();
        let m = grid.len();
        let n = grid[0].len();
        for j in 0..m {
            let mut st = 0;
            for i in 0..n {
                st |= grid[j][i] << i;
            }
            mp.insert(st, j as i32);
        }
        if mp.contains_key(&0) {
            ans.push(*mp.get(&0).unwrap());
            return ans;
        }
        for (&x, &i) in &mp {
            for (&y, &j) in &mp {
                if x & y == 0 {
                    return vec![i.min(j), i.max(j)];
                }
            }
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::good_subsetof_binary_matrix(vec![
        vec![0,1,1,0],
        vec![0,0,0,1],
        vec![1,1,1,1],
    ]))
}