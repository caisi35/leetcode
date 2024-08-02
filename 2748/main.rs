struct Solution{}


impl Solution {
    pub fn count_beautiful_pairs(nums: Vec<i32>) -> i32 {
        let mut res = 0;
        let n = nums.len() as i32;
        for i in 0..n {
            let mut x = nums[i as usize];
            while x >= 10 {
                x /= 10;
            }
            for j in (i+1)..n {
                if Self::gcd(x, nums[j as usize] %10) {
                    res += 1;
                }
            }
        }
        res
    }

    fn gcd(mut a: i32, mut b: i32) -> bool {
        while b != 0 {
            let tmp = b;
            b = a % b ;
            a = tmp;
        }

        a == 1
    }
}

fn main() {
    println!("{}", Solution::count_beautiful_pairs(vec![2, 5, 1, 4]))
}
