struct Solution {}

impl Solution {
    pub fn min_refuel_stops(target: i32, start_fuel: i32, stations: Vec<Vec<i32>>) -> i32 {
        let n = stations.len();
        let mut dp = vec![0; n+1];
        dp[0] = start_fuel;
        // for i in 0..n {
        //     for j in (0..=i).rev() {
        //         if dp[j] >= stations[i][0] {
        //             dp[j+1] = std::cmp::max(dp[j+1], dp[j]+stations[i][1]);
        //         }
        //     }
        // }
        for (i, s) in stations.iter().enumerate() {
            for j in (0..=i).rev() {
                if dp[j] >= s[0] {
                    dp[j+1] = std::cmp::max(dp[j+1], dp[j]+s[1]);
                }
            }
        }
        for (i, v) in dp.iter().enumerate() {
            if *v >= target {
                return i as i32
            }
        }
        return -1
    }
}

fn main() {
    println!("{}", Solution::min_refuel_stops(1, 1, vec![vec![0,0]]));
    println!("{}", Solution::min_refuel_stops(100, 10, vec![vec![10,60], vec![20,30], vec![30,30], vec![60, 40]]));
}