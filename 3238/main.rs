struct Solution{}

impl Solution {
    pub fn winning_player_count(n: i32, pick: Vec<Vec<i32>>) -> i32 {
        let mut cnt = vec![vec![0;11]; n as usize];
        for p in pick {
            cnt[p[0] as usize][p[1] as usize] += 1;
        }
        (0..n as usize).filter(|&i|{(0..=10).any(|j| cnt[i][j]>i)}).count() as i32
    }
}

fn main() {
    println!("{}", Solution::winning_player_count(5, vec![vec![1,1], vec![2,4], vec![2,4], vec![2,4]]));
}