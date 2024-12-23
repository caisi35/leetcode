struct Solution {}

impl Solution {
    pub fn satisfies_conditions(grid: Vec<Vec<i32>>) -> bool {
        let n = grid.len();
        let m = grid[0].len();
        for i in 0..n {
            for j in 0..m {
                if ((i < n-1) && !(grid[i][j] == grid[i+1][j])) || ((j < m-1) && !(grid[i][j] != grid[i][j+1])) {
                    return false;
                }
            }
        }
        return true;
    }
}

fn main() {
    println!("{}", Solution::satisfies_conditions(vec![
        vec![1],
        vec![2],
        vec![3],
    ]))
}

