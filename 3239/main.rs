struct Solution {}

impl Solution {
    pub fn min_flips(grid: Vec<Vec<i32>>) -> i32 {
        let mut row_cnt = 0;
        let mut col_cnt = 0;
        let m = grid.len();
        let n = grid[0].len();
        for i in 0..m {
            for j1 in 0..n/2 {
                let j2 = n-1-j1;
                if grid[i][j1] != grid[i][j2] {
                    row_cnt += 1;
                }
            }
        }
        for j in 0..n {
            for i1 in 0..m/2 {
                let i2 = m-1-i1;
                if grid[i1][j] != grid[i2][j] {
                    col_cnt += 1;
                }
            }
        }
        row_cnt.min(col_cnt)
    }
}

fn main() {
    println!("{}", Solution::min_flips(vec![vec![1,0,0], vec![0,0,0], vec![0,0,1]]));
}