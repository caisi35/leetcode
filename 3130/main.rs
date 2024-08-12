struct Solution{} 

impl Solution {
    pub fn number_of_stable_arrays(zero: i32, one: i32, limit: i32) -> i32 {
        const MOD: i32 = 1000000007;
        let mut memo = vec![vec![vec![-1;2]; (one+1) as usize]; (zero+1) as usize];
        fn dp(zero: usize, one: usize, last_bit: usize, limit: usize, memo: &mut Vec<Vec<Vec<i32>>>) -> i32 {
            if zero == 0 {
                return if last_bit == 0 || one > limit { 0 } else { 1 };
            } else if one == 0 {
                return if last_bit == 1 || zero > limit { 0 } else { 1 };
            }
            if memo[zero][one][last_bit] == -1 {
                let mut res = 0;
                if last_bit == 0 {
                    res = (dp(zero - 1, one, 0, limit, memo) + dp(zero - 1, one, 1, limit, memo)) % MOD;
                    if zero > limit {
                        res = (res - dp(zero - limit - 1, one, 1, limit, memo) + MOD) % MOD;
                    }
                } else {
                    res = (dp(zero, one-1, 0, limit, memo) + dp(zero, one-1, 1, limit, memo)) % MOD;
                    if one > limit {
                        res = (res - dp(zero, one-limit-1, 0, limit, memo) + MOD) % MOD;
                    }
                }
                memo[zero][one][last_bit] = res % MOD;
            }
            memo[zero][one][last_bit]
        }
        let zero = zero as usize;
        let one = one as usize;
        let limit = limit as usize;
        (dp(zero, one, 0, limit, &mut memo) + dp(zero, one, 1 ,limit, &mut memo)) % MOD
    }
}

fn main() {
    println!("{}", Solution::number_of_stable_arrays(1, 1, 2))
}