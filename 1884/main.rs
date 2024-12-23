struct Solution {}

impl Solution {
    pub fn two_egg_drop(n: i32) -> i32 {
        let mut f = vec![i32::MAX/2; (n+1) as usize];
        f[0] = 0;
        for i in 1..=n {
            for k in 1..=i {
                f[i as usize] = std::cmp::min(f[i as usize], std::cmp::max(f[(i-k)as usize], k as i32-1)+1);
            }
        }
        f[n as usize]
    }
}

fn main() {
    println!("{}", Solution::two_egg_drop(100));
}