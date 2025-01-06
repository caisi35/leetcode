
struct ATM {
    cnt: Vec<i64>,
    value: Vec<i64>,
}

impl ATM {
    fn new() -> Self {
        ATM {
            cnt: vec![0,0,0,0,0],
            value: vec![20, 50, 100, 200, 500],
        }
    }

    fn deposit(&mut self, banknotes_count: Vec<i32>) {
        for i in 0..5 {
            self.cnt[i] += banknotes_count[i] as i64;
        }
    }

    fn withdraw(&mut self, mut amount: i32) -> Vec<i32> {
        let mut res = vec![0; 5];
        for i in (0..5).rev() {
            res[i] = std::cmp::min(self.cnt[i], amount as i64 / self.value[i]) as i32;
            amount -= res[i] * self.value[i] as i32;
        }
        if amount > 0 {
            vec![-1]
        } else {
            for i in 0..5 {
                self.cnt[i] -= res[i] as i64;
            }
            res
        }
    }
}

fn main() {
    let mut obj = ATM::new();
    obj.deposit(vec![0,0,1,2,1]);
    println!("{:#?}", obj.withdraw(600));
}