#[warn(unused_comparisons)]

struct Solution {}


impl Solution {
    pub fn find_integers(n: i32) -> i32 {
        let mut ans = 0;
        let mut dp = vec![0; 31];
        dp[0] = 1;
        dp[1] = 1;
        // println!("{:#?}", dp);
        for i in 2..31 {
            dp[i] = dp[i-1] + dp[i-2]
        }
        let mut i: i32 = 29;
        let mut pre = 0;
        while i >= 0 {
            let val = 1 << i;
            if n & val > 0 {
                ans += dp[i as usize +1];
                if pre == 1 {
                    break;
                }
                pre = 1;
            } else {
                pre = 0;
            }
            if i == 0 {
                ans += 1;
            }
            i -= 1;
        }
        ans as i32
    }
}

fn main() {
    println!("{}", Solution::find_integers(5));
}