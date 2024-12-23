use std::collections::HashMap;
use std::cmp::min;

struct Solution{}

impl Solution {
    pub fn mincost_tickets(days: Vec<i32>, mut costs: Vec<i32>) -> i32 {
        let mut memo = vec![0; 366];
        let mut day_m: HashMap<i32, bool> = HashMap::new();
        for d in days {
            day_m.insert(d, true);
        }
        fn dfs (day: i32,  memo:&mut  Vec<i32>, day_m:&mut  HashMap<i32, bool>, costs:&mut  Vec<i32>) -> i32 {
            if day > 365 {
                return 0;
            }
            if memo[day as usize] > 0 {
                return memo[day as usize];
            }
            if day_m.get(&day).is_some() {
                memo[day as usize] = min(min(dfs(day+1, memo, day_m, costs)+costs[0], dfs(day+7, memo, day_m, costs)+costs[1]), dfs(day+30, memo, day_m, costs)+costs[2]);
            } else {
                memo[day as usize] = dfs(day+1, memo, day_m, costs);
            }
            return memo[day as usize];
        }
        return dfs(1,&mut  memo, &mut  day_m, &mut costs);
    }
}

fn main() {
    println!("{}", Solution::mincost_tickets(vec![1,4,6,7,8,20], vec![2,7,15]));

}