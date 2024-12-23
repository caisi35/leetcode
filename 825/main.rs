struct Solution {}

impl Solution {
    pub fn num_friend_requests(ages: Vec<i32>) -> i32 {
        let mut left = 0;
        let mut right = 0;
        let mut ans = 0;
        let mut ages = ages;
        ages.sort();
        for &age in ages.iter() {
            if age < 15 {
                continue
            }
            while ages[left] * 2 <= age + 14 {
                left += 1;
            }
            while right+1 < ages.len() && ages[right+1] <= age {
                right += 1;
            }
            ans += right-left
        }
        return ans as i32;
    }
}

fn main() {
    println!("{}", Solution::num_friend_requests(vec![20,30,100,110,120]));  // 3
}