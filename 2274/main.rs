struct Solution {}

impl Solution {
    pub fn max_consecutive(bottom: i32, top: i32, special: Vec<i32>) -> i32 {
        let mut special = special;
        special.push(bottom - 1);
        special.push(top + 1);
        special.sort();
        let mut ans = 0;
        for i in 0..special.len() - 1 {
            ans = ans.max(special[i+1] - special[i] - 1)
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::max_consecutive(3, 15, vec![7,9,13])); 
}