struct Solution{}

impl Solution {
    pub fn maxmium_score(mut cards: Vec<i32>, cnt: i32) -> i32 {
        cards.sort_by(|a,b| b.cmp(a));
        let mut tmp = 0;
        let mut ans = 0;
        let mut odd = -1;
        let mut even = -1;
        for i in 0..cnt {
            tmp += cards[i as usize];
            if cards[i as usize] % 2 == 0 {
                even = cards[i as usize];
            } else {
                odd = cards[i as usize];
            }
        }

        if tmp % 2 == 0 {
            return tmp
        }

        for i in cnt as usize..cards.len() {
            if cards[i as usize] % 2 == 0 {
                if odd != -1 {
                    ans = std::cmp::max(ans, tmp - odd + cards[i as usize]);
                }
            } else {
                if even != -1 {
                    ans = std::cmp::max(ans, tmp - even + cards[i as usize]);
                }
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::maxmium_score(vec![9,5,9,1,6,10,3,4,5,1], 2))
}