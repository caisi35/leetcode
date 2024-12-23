struct Solution {}

impl Solution {
    pub fn count_k_constraint_substrings(s: String, k: i32, queries: Vec<Vec<i32>>) -> Vec<i64> {
        let n = s.len();
        let s: Vec<u8> = s.bytes().collect();
        let mut count = [0, 0];
        let mut right: Vec<usize> = vec![n; n];
        let mut prefix: Vec<i64> = vec![0; n+1];
        let mut i = 0;
        for j in 0..n {
            count[s[j] as usize - b'0' as usize] += 1;
            while count[0] > k && count[1] > k {
                count[s[i] as usize - b'0' as usize] -= 1;
                right[i] = j;
                i += 1;
            }
            prefix[j+1] = prefix[j] + (j-i+1) as i64;
        }
        let mut res: Vec<i64> = Vec::with_capacity(queries.len());
        for query in queries {
            let l = query[0] as usize;
            let r = query[1] as usize;
            let i = std::cmp::min(right[l], r+1);
            let part1 = prefix[r+1] - prefix[i];
            let part2 = (i-l+1) * (i-l) / 2;
            res.push(part1 as i64 + part2 as i64);
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::count_k_constraint_substrings("010101".to_string(), 1, vec![
        vec![0,5],
        vec![1,4],
        vec![2,3],
    ]));
}