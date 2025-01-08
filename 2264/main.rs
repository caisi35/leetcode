struct Solution {}

impl Solution {
    pub fn largest_good_integer(num: String) -> String {
        let mut res = String::new();
        let n = num.len();
        for i in 0..n-2 {
            if num[i..i+1] == num[i+1..i+2] && num[i+1..i+2] == num[i+2..i+3] {
                let cur = &num[i..i+3];
                if res.is_empty() || cur > &res {
                    res = cur.to_string();
                }
            }
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::largest_good_integer("67771333339".to_string()));
}