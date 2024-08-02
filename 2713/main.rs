use std::collections::HashMap;

struct Solution{}

impl Solution {
    pub fn max_increasing_cells(mat: Vec<Vec<i32>>) -> i32 {
        let m = mat.len();
        let n = mat[0].len();
        let mut mp: HashMap<i32, Vec<(usize, usize)>> = HashMap::new();
        let mut row = vec![0; m];
        let mut col = vec![0; n];

        for i in 0..m {
            for j in 0..n {
                mp.entry(mat[i][j]).or_insert(Vec::new()).push((i, j));
            }
        }

        let mut sorted_map: Vec<_> = mp.iter().collect();
        sorted_map.sort_by(|a, b| a.0.cmp(b.0));
        for (_, pos) in sorted_map {
            let res: Vec<_> = pos.iter().map(|&(i, j) | row[i].max(col[j])+1).collect();
            for (&(i, j), &d) in pos.iter().zip(res.iter()) {
                row[i] = row[i].max(d);
                col[j] = col[j].max(d);
            }
        }
        *row.iter().max().unwrap()
    }
}

fn main() {
    println!("{}", Solution::max_increasing_cells(
        vec![
            vec![3, 1],
            vec![3, 4]
        ]
    ))
}