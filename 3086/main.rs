struct Solution {}

impl Solution {
    pub fn minimum_moves(nums: Vec<i32>, k: i32, max_changes: i32) -> i64 {
        let n = nums.len() as i32;
        let mut index_sum = vec![0 as i64; (n + 1) as usize];
        let mut sum = vec![0 as i64; (n + 1) as usize];
        for i in 0..n {
            index_sum[(i + 1) as usize] = index_sum[i as usize] + (nums[i as usize] as i64) * (i as i64);
            sum[(i + 1) as usize] = sum[i as usize] + (nums[i as usize] as i64);
        }
        let mut res = (((1 as u64) << 63) - 1) as i64;
        for i in 0..n {
            if Self::f(i, &nums) + max_changes >= k {
                if k <= Self::f(i, &nums) {
                    res = std::cmp::min(res, (k - nums[i as usize]) as i64);
                } else {
                    res = std::cmp::min(res, (2 * k - Self::f(i, &nums) - nums[i as usize]) as i64);
                }
                continue;
            }
            let (mut left, mut right) = (0, n);
            while left <= right {
                let mid = (left + right) / 2;
                let (i1, i2) = (std::cmp::max(i - mid, 0), std::cmp::min(i+mid, n-1));
                if sum[(i2 + 1) as usize] - sum[i1 as usize] >= (k - max_changes) as i64 {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            }
            let (mut i1, i2) = (std::cmp::max(i-left, 0), std::cmp::min(i + left, n - 1));
            if sum[(i2 + 1) as usize] - sum[i1 as usize] > (k - max_changes) as i64 {
                i1 += 1;
            }
            let (count1, count2) = (sum[(i + 1) as usize] - sum[i1 as usize], sum[(i2 + 1) as usize] - sum[(i + 1) as usize]);
            res = std::cmp::min(res, index_sum[(i2 + 1) as usize] - index_sum[(i + 1) as usize] - (i as i64) * count2 + (i as i64) * count1- (index_sum[(i + 1) as usize] - index_sum[i1 as usize]) + (2 * max_changes as i64));
        }
        res
    }

    fn f(i: i32, nums: &Vec<i32>) -> i32 {
        let mut x = nums[i as usize];
        if i - 1 >= 0 {
            x += nums[(i - 1) as usize];
        }
        if (i + 1) < nums.len() as i32 {
            x += nums[(i + 1) as usize];
        }
        x
    }
}

fn main() {
    println!("{}", Solution::minimum_moves(vec![0,0,0,0], 2, 3))
}