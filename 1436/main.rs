use std::collections::HashSet;

struct Solution{}

impl Solution {
    pub fn dest_city(paths: Vec<Vec<String>>) -> String {
        let mut c: HashSet<String> = HashSet::new();
        for path in &paths {
            c.insert(path[0].clone());
        }
        for path in &paths {
            if !c.contains(&path[1]) {
                return path[1].clone();
            }
        }
        String::from("")
    }
}

fn main() {
    println!("{}", Solution::dest_city(vec![
        vec!["B".to_string(), "C".to_string()],
        vec!["D".to_string(), "B".to_string()],
        vec!["C".to_string(), "A".to_string()],
    ]))
}