use std::convert::TryInto;

struct Solution{}

impl Solution {
    pub fn max_div_score(nums: Vec<i32>, divisors: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut cnt = -1;
        for &divisor in divisors.iter() {
            let mut tmp = 0;
            for num in nums.iter() {
                if num % divisor == 0 {
                    tmp += 1;
                }
            }
            if tmp > cnt || tmp == cnt && divisor < ans {
                cnt = tmp;
                ans = divisor;
            }
        }
        (ans as usize).try_into().unwrap()
    }
}

fn main() {
    println!("{}", Solution::max_div_score(vec![20,14,21,10], vec![5,7,5]))
}