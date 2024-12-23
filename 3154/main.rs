struct Solution {}

impl Solution {
    pub fn ways_to_reach_stair(k: i32) -> i32 {
        let mut n: i32 = 0;
        let mut npow: i32 = 1;
        let mut ans: i32 = 0;
        loop {
            if npow - n - 1 <= k && k <= npow {
                ans += Self::comb(n + 1, npow - k);
            } else if npow - n - 1 > k {
                break;
            }
            n += 1;
            npow *= 2;
        }
        return ans;
    }

    fn comb(n: i32, k: i32) -> i32 {
        let mut ans: i64 = 1;
        let mut i = n;
        while i >= n-k+1 {
            ans *= i as i64;
            ans /= (n-i+1) as i64;
            i -= 1;
        }
        ans as i32
    }
}

fn main() {
    println!("{}", Solution::ways_to_reach_stair(1));    // 4
    println!("{}", Solution::ways_to_reach_stair(536870891));    // 14307150

}