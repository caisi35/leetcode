struct Solution {}

impl Solution {
    pub fn minimum_difference(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        let mut bits = vec![-1; 31];    
        let mut res = i32::MAX;

        for i in 0..n {
            for j in 0..=30 {
                if nums[i] >> j & 1 == 1 {
                    bits[j] = i as i32;
                }
            }
            let mut pos = Vec::new();
            for j in 0..=30 {
                if bits[j] != -1 {
                    pos.push((bits[j], j as i32));
                }
            }
            pos.sort_by(|a, b|b.0.cmp(&a.0));
            let mut val = 0;
            let mut j = 0;
            while j < pos.len() {
                let p = j;
                while j < pos.len() && pos[j].0 == pos[p].0 {
                    val |= 1 << pos[j].1;
                    j += 1;
                }
                res = std::cmp::min(res, (val - k).abs());
            }
        }
        return res
    }
}


fn main() {
    println!("{}", Solution::minimum_difference(vec![1,2,3,4,5], 2));
}