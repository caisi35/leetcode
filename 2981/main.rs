
struct Solution{}

impl Solution {
    pub fn maximum_length(s: String) -> i32 {
        let mut ans = -1;
        let len = s.len();
        let mut chs: Vec<Vec<i32>> = vec![vec![]; 26];
        let mut cnt = 0;
        let s_bytes = s.as_bytes();
        for i in 0..len {
            cnt += 1;
            if i + 1 == len || s_bytes[i] != s_bytes[i+1] {
                let ch = (s_bytes[i] - b'a') as usize;
                chs[ch].push(cnt);
                cnt = 0;
                for j in (1..chs[ch].len()).rev() {
                    if chs[ch][j] > chs[ch][j-1] {
                        chs[ch].swap(j, j-1);
                    } else {
                        break;
                    }
                }
                if chs[ch].len() > 3 {
                    chs[ch].pop();
                }
            }
        }
        for i in 0..26 {
            if chs[i].len() > 0 && chs[i][0] > 2 {
                ans = ans.max(chs[i][0] - 2);
            }
            if chs[i].len() > 1 && chs[i][0] > 1 {
                ans = ans.max((chs[i][0] - 1).min(chs[i][1]));
            }
            if chs[i].len() > 2 {
                ans = ans.max(chs[i][2]);
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::maximum_length("abcaba".to_string()))
}