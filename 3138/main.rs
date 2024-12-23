struct Solution {}

impl Solution {
    pub fn min_anagram_length(s: String) -> i32 {
        let n = s.len();
        let check = |m: usize| -> bool {
            let mut count0 = vec![0; 26];
            for j in (0..n).step_by(m) {
                let mut c1 = vec![0; 26];
                for k in j..j+m {
                    c1[s.as_bytes()[k] as usize - b'a' as usize] += 1;
                }
                if j > 0 && count0 !=c1 {
                    return false;
                }
                count0 = c1;
            }
            true
        };
        for i in 1..n {
            if n % i != 0 {
                continue
            }
            if check(i) {
                return i as i32;
            }
        }
        n as i32
    }
}

fn main() {
    println!("{}", Solution::min_anagram_lenght("abba".to_string()));
}