struct Solution {}

impl Solution {
    pub fn number_of_alternating_groups(colors: Vec<i32>, k: i32) -> i32 {
        let n = colors.len() as i32;
        let mut res = 0;
        let mut cnt = 1;
        for i in (-k + 2)..n {
            if colors[((i+n) % n) as usize] != colors[((i-1+n)%n) as usize] {
                cnt+=1;
            } else {
                cnt = 1;
            }
            if cnt >= k {
                res += 1;
            }
        }
        return res
    }
}

fn main() {
    println!("{:#?}", Solution::number_of_alternating_groups(vec![1,1,0,1], 4));
}