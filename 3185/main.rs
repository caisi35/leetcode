struct Solution {}

impl Solution {
    pub fn count_complete_day_pairs(hours: Vec<i32>) -> i64 {
        let mut ans: i64 = 0;
        let mut m = vec![0;24];
        for h in hours.iter() {
            ans += m[((24-h%24)%24) as usize] as i64;
            m[(h%24) as usize] += 1;
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::count_complete_day_pairs(vec![12,12,48,22,24]));
}