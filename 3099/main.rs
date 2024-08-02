struct Solution{}

impl Solution {
    pub fn sum_of_the_digits_of_harshad_number(x: i32) -> i32 {
        let mut ans = -1;
        let mut t = 0;
        let mut y = x;
        while y != 0 {
            t += y % 10;
            y /= 10;
        }
        if x % t == 0 {
            ans = t;
        }
        return ans
    }
}

fn main() {
    println!("{}", Solution::sum_of_the_digits_of_harshad_number(18));
    println!("{}", Solution::sum_of_the_digits_of_harshad_number(23));
}