
struct Solution {}

impl Solution {
    pub fn missing_rolls(rolls: Vec<i32>, mean: i32, n: i32) -> Vec<i32> {
        let mut miss_sum = (n + rolls.len() as i32) * mean;
        for roll in rolls.iter() {
            miss_sum -= roll
        }

        if miss_sum < n || miss_sum > n * 6 {
            return vec![]
        }

        let q = miss_sum / n;
        let r = miss_sum % n;
        let mut ans = vec![];
        for i in 0..n {
            let mut t = q;
            if i < r {
                t = t + 1
            }
            ans.push(t);

        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::missing_rolls(vec![3,2,4,3], 4, 2))
}