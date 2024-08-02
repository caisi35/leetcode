struct Solution{}

impl Solution {
    pub fn special_perm(nums: Vec<i32>) -> i32 {
        const MOD: i64 = 1_000_000_007;
        let n = nums.len();
        let mut f = vec![vec![-1 as i64; n]; 1 << n];
        fn dfs(state: usize, i:usize, nums:&Vec<i32>, f: &mut Vec<Vec<i64>>) -> i64 {
            let n = nums.len();
            if f[state][i] != -1 {
                return f[state][i];
            }
            if state == (1 << i) {
                return 1;
            }
            f[state][i] = 0;
            for j in 0..n {
                if i == j || state >> j & 1 == 0 {
                    continue;
                }
                if nums[i] % nums[j] != 0 && nums[j] % nums[i] != 0 {
                    continue;
                }
                f[state][i] = (f[state][i] + dfs(state^(1<<i), j, nums, f)) % MOD;
            }
            f[state][i]
        }
        let mut res = 0 as i64;
        for i in 0..n {
            res = (res + dfs((1<<n)-1, i, &nums, &mut f)) % MOD;
        }
        res as i32
    }
}


fn main() {
    println!("{}", Solution::special_perm(vec![2,3,6]))
}