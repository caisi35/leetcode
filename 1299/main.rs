struct Solution {}

impl Solution {
    pub fn replace_elements(arr: Vec<i32>) -> Vec<i32> {
        let n = arr.len();
        let mut ans = vec![0; n];
        ans[n-1] = -1;
        for i in (0..n-1).rev() {
            ans[i] = ans[i+1].max(arr[i+1]);
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::replace_elements(vec![17,18,5,4,6,1]));
}