struct Solution{}

impl Solution {
    pub fn max_distance(arrays: Vec<Vec<i32>>) -> i32 {
        let mut res = 0;
        let mut min_val = arrays[0][0];
        let mut max_val = arrays[0][arrays[0].len() - 1];
        for i in 1..arrays.len() {
            let n = arrays[i].len();
            res = res.max((arrays[i][n-1] - min_val).abs());
            res = res.max((max_val - arrays[i][0]).abs());
            min_val = min_val.min(arrays[i][0]);
            max_val = max_val.max(arrays[i][n-1]);
        }
        res
    }
}

fn main() {
    println!("{}", Solution::max_disance(vec![vec![1], vec![1]]));
}