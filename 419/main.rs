struct Solution{}

impl Solution {
    pub fn count_battleships(board: Vec<Vec<char>>) -> i32 {
        let mut map = board;
        let n: usize = map.len();
        let m = map[0].len();
        let mut ans = 0;
        for i in 0..n {
            for j in 0..m {
                if map[i][j] == 'X' {
                    map[i][j] = '.';
                    let mut k: usize = j + 1;
                    while k < m && map[i][k] == 'X' {
                        map[i][k] = '.';
                        k += 1;
                    }
                    let mut k2 = i + 1;
                    while k2 < n && map[k2][j] == 'X' {
                        map[k2][j] = '.';
                        k2 += 1;
                    }
                    ans += 1;
                }
            }
        }
        ans
    }
}


fn main() {
    println!("{}", Solution::count_battleships(vec![vec!['.']]));
    println!("{}", Solution::count_battleships(vec![vec!['X','X','X']]));
    println!("{}", Solution::count_battleships(vec![vec!['X','.','.','X'],vec!['.','.','.','X'],vec!['.','.','.','X']]));
}