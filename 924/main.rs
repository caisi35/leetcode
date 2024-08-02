use std::collections::{HashMap, VecDeque};

#[derive(Debug)]
struct Solution {}

impl Solution {
    pub fn min_malware_spread(graph: Vec<Vec<i32>>, initial: Vec<i32>) -> i32 {
        let n = graph.len();
        let mut ids: Vec<i32> = vec![0; n];
        let mut id_to_size: HashMap<i32, i32> = HashMap::new();
        let mut id = 0;
        for i in 0..n {
            if ids[i] == 0 {
                id += 1;
                let mut q = VecDeque::from([i]);
                ids[i] = id;
                let mut size = 1;
                while let Some(u) = q.pop_front() {
                    for v in 0..n {
                        if graph[u][v] == 1 && ids[v] == 0 {
                            size += 1;
                            q.push_back(v);
                            ids[v] = id;
                        }
                    }
                }
                id_to_size.insert(id, size);
            }
        }
        let mut id_to_initials: HashMap<i32, i32> = HashMap::new();
        for &u in initial.iter() {
            *id_to_initials.entry(ids[u as usize]).or_insert(0) += 1;
        }
        let mut ans = n as i32 + 1;
        let mut ans_removed = 0;
        for &u in initial.iter() {
            let removed = if id_to_initials[&ids[u as usize]] == 1 {
                id_to_size[&ids[u as usize]]
            } else {
                0
            };
            if removed > ans_removed || (removed == ans_removed && u < ans) {
                ans = u;
                ans_removed = removed;
            }
        }
        ans
    }
}

fn main() {
    println!("main");
    let initial = vec![1, 2];
    let g = vec![vec![1, 1, 1], vec![1, 1, 1], vec![1, 1, 1]];
    println!("{}", Solution::min_malware_spread(g, initial));
}
