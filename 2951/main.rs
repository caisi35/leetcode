
struct Solution {}

impl Solution {
    pub fn find_peaks(mountain: Vec<i32>) -> Vec<i32> {
        let mut ans: Vec<i32> = Vec::new();
        for m in 1..mountain.len()-1 {
            if mountain[m] > mountain[m-1] && mountain[m] > mountain[m+1] {
                ans.push(m as i32)
            }
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::find_peaks(vec![1, 4, 3, 8, 5]))
}

