struct Solution {}

impl Solution {
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        let mut ind = vec![0; n as usize + 1];
        let mut out = vec![0; n as usize + 1];
        for t in trust.iter() {
            ind[t[1] as usize] = ind[t[1] as usize] + 1;
            out[t[0] as usize] = out[t[0] as usize] + 1;
        }
        for i in 1..n+1 {
            if ind[i as usize] == n-1 && out[i as usize] == 0 {
                return i
            }
        }
        -1
    }
}

fn main() {
    println!("{}", Solution::find_judge(3, vec![vec![1,2], vec![2,3]]));
    println!("{}", Solution::find_judge(2, vec![vec![1,2]]));
}