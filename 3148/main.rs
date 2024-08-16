
struct Solution{}

impl Solution {
    pub fn max_score(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut premin = vec![vec![i32::MAX; n]; 2];
        let mut ans = i32::MIN;
        for i in 0..m {
            for j in 0..n {
                let mut pre = i32::MAX;
                if i > 0 {
                    pre = pre.min(premin[(i-1)&1][j]);
                }
                if j > 0 {
                    pre = pre.min(premin[i&1][j-1]);
                }
                if i + j > 0 {
                    ans = ans.max(grid[i][j] - pre);
                }
                premin[i&1][j] = pre.min(grid[i][j]);
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::max_score(vec![
        vec![9,5,7,3],
        vec![8,9,6,1],
        vec![6,7,14,3],
        vec![2,5,3,1],
    ]))
}