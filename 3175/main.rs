struct Solution {}

impl Solution {
    pub fn find_winning_player(skills: Vec<i32>, k: i32) -> i32 {
        let mut i = 0;       
        let mut lasti = 0;       
        let mut cnt = 0;
        let n = skills.len();
        while i < n {
            let mut j = i + 1;
            while j < n && skills[j as usize] < skills[i as usize] && cnt < k {
                cnt += 1;
                j += 1;
            }
            if cnt == k {
                return i as i32
            }
            cnt = 1;
            lasti = i;
            i = j;
        }
        lasti as i32
    }
}

fn main() {
    println!("{}", Solution::find_winning_player(vec![4,2,6,3,9], 2))
}