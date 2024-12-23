struct Solution {}

impl Solution {
    pub fn max_height_of_triangle(red: i32, blue: i32) -> i32 {
        fn h(x: i32, y: i32) -> i32 {
            let odd = 2 * ((x as f64).sqrt() as i32) - 1;
            let even = 2 * (((-1.0 + (1.0 + 4.0 * (y as f64)).sqrt())) / 2.0) as i32;
            odd.min(even) + 1
        }
        std::cmp::max(h(red, blue), h(blue, red))
    }
}

fn main() {
    println!("{}", Solution::max_height_of_triangle(10, 1));
}