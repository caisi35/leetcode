struct Solution{}

impl Solution {
    pub fn is_balanced(num: String) -> bool {
        let mut df = 0;
        let mut s = 1;
        for c in num.chars() {
            let d = c.to_digit(10).unwrap() as i32;
            df += d * s;
            s = -s;
        }
        df == 0
    }
}

fn main() {
    println!("{:#?}", Solution::is_balanced("1234".to_string()));
}