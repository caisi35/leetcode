struct Solution {}

impl Solution{
    pub fn square_is_white(coordinates: String) -> bool {
        

        let c: Vec<char> = coordinates.chars().collect();
        let c1 = 'a' as i32 - c[0] as i32 + 1;
        let c2 = c[1] as i32;
        (c1 + c2) % 2 != 0 
    }
}

fn main() {
    println!("{}", Solution::square_is_white("h3".to_string()));
}