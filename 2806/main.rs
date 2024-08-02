
struct Solution{}

impl Solution {
    pub fn account_balance_after_purchase(purchase_amount: i32) -> i32 {
        let m = purchase_amount % 10;
        let mut s = 0;
        if m >= 5 {
            s = 10;
        }

        return 100 - (purchase_amount - m + s)
    }
}

fn main() {
    println!("{}", Solution::account_balance_after_purchase(9));
    println!("{}", Solution::account_balance_after_purchase(15));
    println!("{}", Solution::account_balance_after_purchase(14));
}