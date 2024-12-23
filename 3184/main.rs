struct Solution {}

impl Solution {
    pub fn count_complete_day_pairs(hours: Vec<i32>) -> i32 {
        let n = hours.len();
        let mut ans = 0;
        for i in 0..n {
            for j in i+1..n {
                if (hours[i] + hours[j]) % 24 == 0 {
                    ans+=1;
                }
            }
        }
        return ans;
    }
}

fn main() {
    println!("{}", Solution::count_complete_day_pairs(vec![12,12,30,24,24]));
}