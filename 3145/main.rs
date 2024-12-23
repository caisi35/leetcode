struct Solution {}

impl Solution {
    pub fn find_products_of_elements(queries: Vec<Vec<i64>>) -> Vec<i32> {
        let mut ans = Vec::new();
        for query in queries.iter() {
            let mut query = query.clone();
            query[0] += 1;
            query[1] += 1;
            let l = Self::mid_check(query[0]);
            let r = Self::mid_check(query[1]);
            let mod_ = query[2];
            let mut res: i64 = 1;
            let mut pre = Self::count_one(l - 1);
            for j in 0..60 {
                if (1i64 << j) & l != 0 {
                    pre += 1;
                    if pre >= query[0] && pre <= query[1] {
                        res = res * (1i64 << j) % mod_;
                    }
                }
            }

            if r > l {
                let mut bac = Self::count_one(r - 1);
                for j in 0..60 {
                    if (1i64 << j) & r != 0 {
                        bac += 1;
                        if bac >= query[0] && bac <= query[1] {
                            res = res * (1i64 << j) % mod_;
                        }
                    }
                }
            }

            if r - l > 1 {
                let xs = Self::count_pow(r-1) - Self::count_pow(l);
                res = res * Self::pow_mod(2, xs, mod_) % mod_;
            }
            ans.push(res as i32);
        }
        ans
    }

    fn count_one(x: i64) -> i64 {
        let mut res: i64 = 0;
        let mut sum: i64 = 0;
        for i in (0..=60).rev() {
            if (1i64 << i) & x != 0 {
                res += sum * (1i64 << i);
                sum += 1;
                if i > 0 {
                    res += i * (1i64 << (i - 1));
                }
            }
        }
        res += sum;
        res
    }

    fn count_pow(x: i64) -> i64 {
        let mut res: i64  = 0;
        let mut sum: i64 = 0;
        for i in (0..=60).rev() {
            if (1i64 << i) & x != 0 {
                res += sum * (1i64 << i);
                sum += i;
                if i > 0 {
                    res += i * (i - 1) / 2 * (1i64 << (i - 1))
                }
            }
        }
        res += sum;
        res
    }

    fn pow_mod(mut x: i64, mut y: i64, mod_: i64) -> i64 {
        let mut res: i64 = 1;
        while y != 0 {
            if y & 1 != 0 {
                res = res * x % mod_;
            }
            x = x * x % mod_;
            y >>= 1;
        }
        res
    }

    fn mid_check(x: i64) -> i64 {
        let mut l: i64 = 1;
        let mut r: i64 = 1_000_000_000_000_000;
        while l < r {
            let mid = (l + r) >> 1;
            if Self::count_one(mid) >= x {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        r
    }
}

fn main() {
    println!("{:#?}", Solution::find_products_of_elements(vec![vec![1,3,7]]))      // [4]
}