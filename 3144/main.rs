const INF: i32 = 0x3f3f3f3f;

struct Solution {}

impl Solution {
    pub fn minimum_substrings_in_partition(s: String) -> i32 {
        let n = s.len();
        let mut d = vec![INF; n+1];
        d[0] = 0;
        for i in 1..=n {
            let mut max_cnt = 0;
            let mut occ_cnt = std::collections::HashMap::new();
            for j in (1..=i).rev() {
                let c = s.as_bytes()[j-1];
                *occ_cnt.entry(c).or_insert(0) += 1;
                max_cnt = max_cnt.max(*occ_cnt.get(&c).unwrap());
                if max_cnt * occ_cnt.len() == i-j+1 && d[j-1] != INF {
                    d[i] = d[i].min(d[j-1]+1);
                }
            }
        }
        d[n]
    }
}

fn main() {
    println!("{}", Solution::minimum_substrings_in_partition("fabccddg".to_string()));  // 3
}