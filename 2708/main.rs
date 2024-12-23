struct Solution {}

impl Solution {
    pub fn max_strength(nums: Vec<i32>) -> i64 {
        let (mut n, mut z, mut p) = (0, 0, 0);
        let mut prod = 1;
        let mut mn = -100000000;
        for &num in nums.iter() {
            if num < 0 {
                n += 1;
                prod *= num;
                if num > mn {
                    mn = num;
                }
            } else if num == 0 {
                z += 1;
            } else {
                prod *= num;
                p += 1;
            }
        }
        if n == 1 && z == 0 && p == 0 {
            return nums[0] as i64
        }
        if n <= 1 && p == 0 {
            return 0 as i64
        }
        if prod < 0 {
            return (prod / mn) as i64
        } else {
            return prod as i64
        }
    }
}

fn main() {
    println!("{}", Solution::max_strength(vec![-4,-5,-4])) // 20
}