struct Solution{}

impl Solution {
    pub fn minimum_levels(possible: Vec<i32>) -> i32 {
        let n = possible.len() as i32;
        let mut tot = 0;
        for v in possible.iter() {
            tot += v;
        }
        tot = tot * 2 - n;
        let mut pre = 0;
        for i in 0..n-1 {
            if possible[i as usize] == 1 {
                pre += 1;
            } else {
                pre -= 1;
            }
            if 2 * pre > tot {
                return i + 1
            }
        }
        -1
    }
}

fn main() {
    println!("{}", Solution::minimum_levels(vec![1,0,1,0]))
}