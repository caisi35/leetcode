struct Solution {}

impl Solution {
    pub fn num_rook_captures(board: Vec<Vec<char>>) -> i32 {
        let mut cnt = 0;
        let mut st = 0;
        let mut ed = 0;
        let dx = [0,1,0,-1];
        let dy = [1,0,-1,0];
        for i in 0..8 {
            for j in 0..8 {
                if board[i][j] == 'R' {
                    st = i;
                    ed = j;
                    break;
                }
            }
        }
        for i in 0..4 {
            let mut step = 0;
            loop {
                let tx = st as i32 + step * dx[i] as i32;
                let ty = ed as i32 + step * dy[i] as i32;
                if tx < 0 || tx >= 8 || ty < 0 || ty >= 8 || board[tx as usize][ty as usize] == 'B' {
                    break;
                }
                if board[tx as usize][ty as usize] == 'p' {
                    cnt += 1;
                    break;
                }
                step += 1;
            }
        }
        cnt
    }
}

fn main() {
    println!("{}", Solution::num_rook_captures(vec![
        vec!['.', '.','.','.','.','.','.','.'],
        vec!['.', '.','.','p','.','.','.','.'],
        vec!['.', '.','.','R','.','.','.','p'],
        vec!['.', '.','.','.','.','.','.','.'],
        vec!['.', '.','.','.','.','.','.','.'],
        vec!['.', '.','.','p','.','.','.','.'],
        vec!['.', '.','.','.','.','.','.','.'],
        vec!['.', '.','.','.','.','.','.','.'],
    ]))
}