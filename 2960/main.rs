
struct Solution{}

impl Solution {
    pub fn count_tested_devices(battery_percentages: Vec<i32>) -> i32 {
        let mut ans = 0;
        for v in battery_percentages {
            if v > 0 && v - ans > 0 {
                ans += 1
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::count_tested_devices(vec![1, 1, 2, 1, 3]))
}