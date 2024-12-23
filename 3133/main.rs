struct Solution{}

impl Solution {
    pub fn min_end(n: i32, x: i32) -> i64 {
        let bit_count = 64 - n.leading_zeros() - x.leading_zeros();
        let mut res = x as i64;
        let mut j = 0;
        let n = (n-1) as i64;
        for i in 0..bit_count {
            if ((res >> i) & 1) == 0 {
                if ((n >> j) & 1) != 0 {
                    res |= 1 << i;
                }
                j += 1;
            }
        }
        res
    }
}

fn main() {
    println!("{}", Solution::min_end(3, 4)) // 6
}