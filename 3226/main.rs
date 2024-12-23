struct Solution {}

impl Solution {
    pub fn min_changes(mut n: i32, mut k: i32) -> i32 {
        let mut res = 0;
        while n > 0 || k > 0 {
            if (n & 1) == 0 && (k & 1) == 1 {
                return -1;
            }
            if (n & 1) == 1 && (k & 1) == 0 {
                res += 1;
            }
            n >>= 1;
            k >>= 1;
        }
        res
    }
}

fn main() {
    println!("{}", Solution::min_changes(14, 13));
}