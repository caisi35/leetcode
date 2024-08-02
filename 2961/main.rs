struct Solution{}

impl Solution {
    pub fn get_good_indices(veriables: Vec<Vec<i32>>, target: i32) -> Vec<i32> {
        let mut ans: Vec<i32> = Vec::new();
        for i in 0..veriables.len() {
            if Self::pow_mod(Self::pow_mod(veriables[i][0], veriables[i][1], 10), veriables[i][2], veriables[i][3]) == target {
                ans.push(i as i32);
            }
        }
        ans
    }

    fn pow_mod(mut x: i32, mut y: i32, m: i32) -> i32 {
        let mut res = 1;
        while y > 0 {
            if (y & 1 ) == 1 {
                res = res * x % m;
            }
            x = x * x % m;
            y >>= 1;
        }
        res
    }
}

fn main() {
    println!("{:?}", Solution::get_good_indices(vec![vec![39,3,1000,1000]], 17))
}