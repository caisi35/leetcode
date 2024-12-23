struct Solution {}

impl Solution {
    pub fn number_of_alternating_groups(colors: Vec<i32>) -> i32 {
        let mut ans = 0;
        let n = colors.len();
        for i in 0..n {
            if colors[i] != colors[(i+n-1)%n] && colors[i] != colors[(i+1)%n] {
                ans += 1;
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::number_of_alternating_groups(vec![0,0,1]));
}