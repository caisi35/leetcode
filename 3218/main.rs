struct Solution {}

impl Solution {
    pub fn minimum_cost(m: i32, n: i32, horizontal_cut: Vec<i32>, vertical_cut: Vec<i32>) -> i32 {
        let mut memo = vec![-1; (m*m*n*n) as usize];
        fn dp(row1: i32, col1: i32, row2: i32, col2: i32, m: i32, n: i32, horizontal_cut: &Vec<i32>, vertical_cut: &Vec<i32>, memo: &mut Vec<i32>) -> i32 {
            let index = |row1:i32, col1:i32, row2:i32, col2:i32| -> usize {
                ((row1 * n + col1) * m * n + row2 * n + col2) as usize
            };
            if row1 == row2 && col1 == col2 {
                return 0;
            }
            let ind = index(row1, col1, row2, col2);
            if memo[ind] >= 0 {
                return memo[ind];
            }
            memo[ind] = i32::MAX;
            for i in row1..row2 {
                memo[ind] = memo[ind].min(dp(row1, col1, i, col2, m, n, horizontal_cut, vertical_cut, memo) + 
                dp(i+1, col1, row2, col2, m, n, horizontal_cut, vertical_cut, memo) +
                horizontal_cut[i as usize]);
            }
            for i in col1..col2 {
                memo[ind] = memo[ind].min(dp(row1, col1, row2, i, m, n, horizontal_cut, vertical_cut, memo) + 
                dp(row1, i+1, row2, col2, m, n, horizontal_cut, vertical_cut, memo) + 
                vertical_cut[i as usize]);
            }
            memo[ind]
        }
        dp(0, 0, m-1, n-1, m, n, &horizontal_cut, &vertical_cut, &mut memo)
    }
}

fn main() {
    println!("{}", Solution::minimum_cost(3, 2, vec![1,3], vec![5]));
}