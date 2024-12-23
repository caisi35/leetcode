struct Solution {}

impl Solution {
    pub fn losing_player(x: i32, y: i32) -> String {
        let num = x.min(y/4);
        if num % 2 == 0 {
            return "Bob".to_string();
        }
        "Alice".to_string()
    }
}

fn main() {
    println!("{}", Solution::losing_player(4, 11));
    println!("{}", Solution::losing_player(2, 7));
}