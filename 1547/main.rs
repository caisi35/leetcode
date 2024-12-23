struct Solution {}

impl Solution {
    pub fn min_cost(n: i32, cuts: Vec<i32>) -> i32 {
        let mut cuts = cuts;
        cuts.push(0);
        cuts.push(n);
        cuts.sort();
        let m = cuts.len();
        let mut f = vec![vec![0; m]; m];
        for i in (0..m).rev() {
            for j in i+2..m {
                f[i][j] = i32::MAX;
                for k in i+1..j {
                    f[i][j] = std::cmp::min(f[i][j], f[i][k] + f[k][j]);
                }
                f[i][j] += cuts[j] - cuts[i];
            }
        }
        f[0][m-1]
    }
}

fn main() {
    println!("{}", Solution::min_cost(9, vec![5,6,1,4,2]));
}