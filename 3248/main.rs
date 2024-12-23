struct Solution {}

impl Solution {
    pub fn final_position_of_snake(n: i32, commands: Vec<String>) -> i32 {
        let mut res = 0;
        for c in commands.iter() {
            if c == "RIGHT" {
                res+=1;
            } else if c == "LEFT" {
                res-=1;
            } else if c == "DOWN" {
                res += n;
            } else {
                res -=n;
            }
        }
        return res;
    }
}

fn main() {
    println!("{}", Solution::final_position_of_snake(2, vec!["RIGHT".to_string(), "DOWN".to_string()]));
}