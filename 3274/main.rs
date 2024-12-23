struct Solution{}

impl Solution {
    pub fn check_two_chessboards(coordinate1: String, coordinate2: String) -> bool {
        (coordinate1.chars().nth(0).unwrap() as i32 - coordinate2.chars().nth(0).unwrap() as i32
        + coordinate1.chars().nth(1).unwrap() as i32 - coordinate2.chars().nth(1).unwrap() as i32) % 2 == 0
    }
}

fn main() {
    println!("{}", Solution::check_two_chessboards("a7".to_string(), "a6".to_string()));
}