struct Solution{}

impl Solution {
    pub fn smallest_string(s: String) -> String {
        let target = 'a';
        let first_a = s.chars().position(|c|c!=target).unwrap_or(s.len());
        if first_a == s.len() {
            return s[0..s.len()-1].to_string() + "z";
        }
        let second_a = Self::find_second_a(&s, first_a);
        let mut res = String::new();
        for (i, c) in s.chars().enumerate() {
            if first_a <= i && i < second_a {
                res.push((c as u8-1) as char);
            } else {
                res.push(c);
            }
        }
        res
    }

    fn find_second_a(s: &String, first_a: usize) -> usize {
        let chars = s.chars().skip(first_a);
        for (index, c) in chars.enumerate() {
            if c == 'a' {
                return index + first_a;
            }
        }
        s.len()
    }
}

fn main() {
    println!("{}", Solution::smallest_string("cbabc".to_string()))
}