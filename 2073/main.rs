struct Solution{}

impl Solution {
    pub fn time_requried_to_buy(tikects: Vec<i32>, k: i32) -> i32 {
        let mut ans = 0;
        let n = tikects.len() as i32;
        for i in 0..n {
            if i <= k {
                ans += std::cmp::min(tikects[i as usize], tikects[k as usize]);
            } else {
                ans += std::cmp::min(tikects[i as usize], tikects[k as usize] - 1);
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::time_requried_to_buy(vec![5,1,1,1], 0))
}