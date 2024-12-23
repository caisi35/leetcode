struct Solution {}

impl Solution {
    pub fn can_alice_win(nums: Vec<i32>) -> bool {
        let mut s1 = 0;
        let mut s2 = 0;
        for &v in nums.iter() {
            if v < 10 {
                s1 += v;
            } else {
                s2 += v;
            }
        }
        return s1 != s2;
    }
}

fn main() {
    println!("{}", Solution::can_alice_win(vec![5,5,5,25]));
}