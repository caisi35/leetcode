use std::convert::TryInto;

struct Solution{}

impl Solution {
    pub fn distribute_candies(candies: i32, num_people: i32) -> Vec<i32> {
        let mut ans = vec![0;num_people.try_into().unwrap()];
        let mut t = 0;
        let mut c = candies;
        while c > 0 {
            for j in 1..num_people+1 {
                let i = j as usize;
                t += 1;
                if c <= t {
                    ans[i-1] = ans[i-1] + c;
                    c = 0;
                    break
                } else {
                    ans[i-1] = ans[i-1] + t;
                    c -= t;
                }
            }
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::distribute_candies(60, 4))  // [15,18,15,12]
}