use std::collections::VecDeque;

struct Solution {}

impl Solution {
    pub fn shortest_distance_after_queries(n: i32, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut neighbors: Vec<Vec<i32>> = vec![Vec::new(); n as usize];
        for i in 0..n-1 {
            neighbors[i as usize].push(i+1);
        }
        let mut res: Vec<i32> = Vec::new();
        for query in queries {
            neighbors[query[0] as usize].push(query[1]);
            res.push(Self::bfs(n as usize, &neighbors));
        }
        res
    }

    fn bfs(n: usize, neighbors: &Vec<Vec<i32>>) -> i32 {
        let mut dist = vec![-1; n];
        let mut q = VecDeque::new();
        q.push_back(0);
        dist[0] = 0;
        while let Some(x) = q.pop_front() {
            for &y in &neighbors[x] {
                if dist[y as usize] >= 0 {
                    continue;
                }
                q.push_back(y as usize);
                dist[y as usize] = dist[x] + 1;
            }
        }
        dist[n-1]
    }
}

fn main() {
    println!("{:#?}", Solution::shortest_distance_after_queries(5, vec![vec![2,4], vec![0,2],vec![0,4]]));
}