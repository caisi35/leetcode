struct Solution {}

impl Solution {
    pub fn judge_square_sum(c: i32) -> bool {
        for a in 0..=((c as f64).sqrt() as i32) {
            let b = c - a * a;
            if b >= 0 {
                let b_ = (b as f64).sqrt();
                if b_.fract() == 0.0 {
                    return true;
                }
            }
        }
        false
    }
}

fn main() {
    println!("{}", Solution::judge_sequare_sum(3));
}