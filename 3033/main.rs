struct Solution {}

impl Solution {
    pub fn modified_matrix(mut matrix: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let n = matrix.len();
        let m = matrix[0].len();
        for j in 0..m {
            let mut zd = -1;
            for i in 0..n {
                zd = std::cmp::max(zd, matrix[i][j])
            }
            for i in 0..n {
                if matrix[i][j] == -1 {
                    matrix[i][j] = zd
                }
            }
        }
        matrix
    }
}

fn main() {
    println!("{:#?}", Solution::modified_matrix(vec![
        vec![1,2,-1],
        vec![4,-1, 6],
        vec![7,8,9]
    ]))
}