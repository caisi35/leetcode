struct Solution {}

impl Solution {
    pub fn get_kth(lo: i32, hi: i32, k: i32) -> i32 {
        let mut v: Vec<i32> = (lo..=hi).collect();
        v.sort_by(|&u, &v| {
            let f1 = Self::get_f(u);
            let f2 = Self::get_f(v);
            if f1 != f2 {
                f1.cmp(&f2)
            } else {
                u.cmp(&v)
            }
        });
        v[k as usize -1]
    }

    fn get_f(x: i32) -> i32 {
        if x == 1 {
            return 0;
        }
        if  x & 1 == 1 {
            Self::get_f(x * 3 + 1) + 1
        } else {
            Self::get_f(x / 2) + 1
        }
    }
}

fn main() {
    println!("{}", Solution::get_kth(1, 1000, 777));
}