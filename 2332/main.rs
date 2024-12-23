struct Solution {}

impl Solution {
    pub fn latest_time_catch_the_bus(
        mut buses: Vec<i32>,
        mut passengers: Vec<i32>,
        capacity: i32,
    ) -> i32 {
        buses.sort();
        passengers.sort();
        let mut pos = 0;
        let mut space = 0;
        for &arrive in &buses {
            space = capacity;
            while space > 0 && pos < passengers.len() && passengers[pos] <= arrive {
                space -= 1;
                pos += 1;
            }
        }
        pos -= 1;
        let mut last_catch_time = if space > 0 { *buses.last().unwrap() } else { passengers[pos]} ;
        while pos >= 0 && passengers[pos as usize] == last_catch_time {
            pos -= 1;
            last_catch_time -= 1;
        }
        return last_catch_time;
    }
}

fn main() {
    println!(
        "{}",
        Solution::latest_time_catch_the_bus(vec![10, 20], vec![2, 17, 18, 19], 2)
    ); // 16
    println!(
        "{}",
        Solution::latest_time_catch_the_bus(vec![2], vec![2, 3], 2)
    ); // 16
}
