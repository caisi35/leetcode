struct Solution {}

impl Solution {
    pub fn sort_the_students(score: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let mut score = score.clone();
        score.sort_by(|u,v| v[k as usize].cmp(&u[k as usize]));
        score
    }
}

fn main() {
    println!("{:#?}", Solution::sort_the_students(vec![vec![3,4], vec![5,6]], 0));
}