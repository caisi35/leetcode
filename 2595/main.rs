struct Solution{}

impl Solution {
    pub fn even_odd_bit(n: i32) -> Vec<i32> {
        let mut res = vec![0,0];
        let mut i = 0;
        let mut n = n;
        while n > 0 {
            res[i as usize] += n & 1;
            n >>= 1;
            i ^= 1;
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::even_odd_bit(50));
}