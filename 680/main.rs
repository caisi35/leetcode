struct Solution{}

impl Solution {
    pub fn valid_palindrome(s: String) -> bool {
        let mut low = 0;
        let mut high = s.len()-1;
        while low < high {
            if s.as_bytes()[low] == s.as_bytes()[high] {
                low+=1;
                high-=1;
            } else {
                return Solution::check(&s, low, high-1) || Solution::check(&s, low+1, high)
            }
        }
        true
    }

    pub fn check(s: &str, low: usize, high: usize) -> bool {
        let s = s.as_bytes();
        let (mut low, mut high) = (low, high);
        while low < high {
            if s[low] != s[high] {
                return false;
            }
            low += 1;
            high -= 1;
        }
        true
    }
}

fn main() {
    println!("{:#?}", Solution::valid_palindrome("abca".to_string()));
}