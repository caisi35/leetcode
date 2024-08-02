struct Solution{}

impl Solution {

    pub fn number_of_right_triangles(grid: Vec<Vec<i32>>) -> i64 {
        let mut ans = 0;
        let mut col: Vec<i32> = vec![0; grid[0].len()];
        for i in 0..grid.len() {
            for j in 0..grid[i].len() {
                col[j] += grid[i][j];
            }
        }

        for i in 0..grid.len() {
            let mut row_1 = 0;
            for j in 0..grid[i].len() {
                row_1 += grid[i][j];
            }
            for j in 0..grid[i].len() {
                if grid[i][j] == 1 {
                    ans += (row_1 as i64 -1) * (col[j] as i64 -1);
                }
            }
        }
        ans as i64
    }
}

fn main() {
    println!("{}", Solution::number_of_right_triangles(vec![
        vec![0,1,0],
        vec![0,1,1],
        vec![0,1,0],
    ]))
}