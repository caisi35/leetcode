struct Solution{}

impl Solution {
    pub fn find_maximum_number(k: i64, x: i32) -> i64 {
        let (mut l, mut r) = (1i64, (k + 1) << x);
        while l < r {
            let m = (l + r + 1) / 2;
            if Self::accumulated_price(x, m) > k {
                r = m - 1;
            } else {
                l = m;
            }
        }
        return l;
    }

    fn accumulated_bit_price(x: i32, num: i64) -> i64 {
        let period = 1i64 << x;
        let mut res = period / 2 * (num / period);
        if num % period >= period / 2 {
            res += num % period - (period / 2 - 1);
        }
        return res;
    }

    fn accumulated_price(x: i32, num: i64) -> i64 {
        let mut res = 0i64;
        let length = 64 - num.leading_zeros();
        for i in (x..=length as i32).step_by(x as usize) {
            res += Self::accumulated_bit_price(i, num);
        }
        return res;
    }
}

fn main() {
    println!("{}", Solution::find_maximum_number(9, 1))  // 6
}