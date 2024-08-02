struct Solution{}

impl Solution {
    pub fn find_value_of_partition(mut nums: Vec<i32>) -> i32 {
        let mut ans = 1000000000;
        let n = nums.len();
        nums.sort();
        for i in 0..n - 1 {
            ans = std::cmp::min((nums[i] - nums[i+1]).abs(), ans);
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::find_value_of_partition(vec![1,3,2,4]));
    println!("{}", Solution::find_value_of_partition(vec![84,11,100,100,75]));

}