struct Solution{}

impl Solution {
    pub fn max_coins(piles: Vec<i32>) -> i32 {
        let mut piles = piles;
        piles.sort();
        let l = piles.len();
        let r = l / 3;
        let mut coins = 0;
        let mut i: i32 = (l - 2) as i32;
        for _ in 0..r {
            coins += piles[i as usize];
            i -= 2;
        }
        coins
    }
}

fn main() {
    println!("{:#?}", Solution::max_coins(vec![2,4,5]));
}