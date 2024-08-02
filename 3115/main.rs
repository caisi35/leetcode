use std::collections::HashSet;

struct Solution {}

impl Solution {
    pub fn maximum_prime_defierence(nums: Vec<i32>) -> i32 {
        let ps = Self::primes();
        let mut f = -1;
        let mut ans = 0;
        for (i, &num) in nums.iter().enumerate() {
            if ps.contains(&num) {
                if f != -1 {
                    ans = ans.max(i as i32 - f);
                } else {
                    f = i as i32;
                }
            }
        }
        ans
    }

    fn primes() -> HashSet<i32> {
        let mut ps: HashSet<i32> = HashSet::new();
        for i in 2..100 {
            let mut f = true;
            for j in 2..i {
                if i % j == 0 {
                    f = false;
                }
            }
            if f {
                ps.insert(i);
            }
        }
        ps
    }
}

fn main() {
    println!("{}", Solution::maximum_prime_defierence(vec![4,2,9,5,3]))
}