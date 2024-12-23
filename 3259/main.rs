struct Solution {}

impl Solution {
    pub fn max_energy_boost(energy_drink_a: Vec<i32>, energy_drink_b: Vec<i32>) -> i64 {
        let n = energy_drink_a.len();
        let mut d = vec![vec![0; 2]; n+ 1];
        for i in 1..=n {
            d[i][0] = d[i-1][0] + energy_drink_a[i-1] as i64;
            d[i][1] = d[i-1][1] + energy_drink_b[i-1] as i64;
            if i >= 2 {
                d[i][0] = d[i][0].max(d[i-2][1] + energy_drink_a[i-1] as i64);
                d[i][1] = d[i][1].max(d[i-2][0] + energy_drink_b[i-1] as i64);
            }
        }
        d[n][0].max(d[n][1])
    }
}

fn main() {
    println!("{}", Solution::max_energy_boost(vec![4,1,1], vec![1,1,3]));
}