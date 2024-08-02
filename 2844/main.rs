struct Solution {}

impl Solution {
    pub fn minimum_operations(num: String) -> i32 {
        let n = num.len();
        let mut f1 = false;
        let mut z = false;
        let num_chars: Vec<char> = num.chars().collect();

        for i in (0..n).rev() {
            let t = num_chars[i];
            if t == '0' || t == '5' {
                if z {
                    return (n - i - 2) as i32;
                } 
                if t == '5' {
                    f1 = true;
                } else if t == '0' {
                    z = true;
                }
            }
    
            if t == '2' || t == '7' {
                if f1 {
                    return (n - i - 2) as i32;
                }
            }
        }
        if z {
            return (n-1) as i32;
        }
        return n as i32;
    }
}

fn main() {
    println!("{}", Solution::minimum_operations("2245047".to_string()));
    println!("{}", Solution::minimum_operations("10".to_string()));
    println!("{}", Solution::minimum_operations("2908305".to_string()));

}