use std::collections::BinaryHeap;
use std::cmp::Reverse;

struct Solution {}

impl Solution {
    pub fn minimum_time(n: i32, edges: Vec<Vec<i32>>, disappear: Vec<i32>) -> Vec<i32> {
        let mut adj = vec![Vec::new(); n as usize];
        for edge in edges.iter() {
            let u = edge[0] as usize;
            let v = edge[1] as usize;
            let length = edge[2];
            adj[u].push((v, length));
            adj[v].push((u, length));
        }

        let mut pq = BinaryHeap::new();
        pq.push(Reverse((0,0)));
        let mut answer = vec![-1; n as usize];
        answer[0] = 0;

        while let Some(Reverse((t, u))) = pq.pop() {
            if t != answer[u] {
                continue;
            }
            for &(v, length) in adj[u].iter() {
                if t + length < disappear[v] && (answer[v] == -1 || t + length < answer[v]) {
                    pq.push(Reverse((t + length, v)));
                    answer[v] = t + length;
                }
            }
        }
        answer
    }
}

fn main() {
    println!("{:#?}", Solution::minimum_time(3, vec![vec![0,1,2],vec![1,2,1],vec![0,2,4]], vec![1,1,5]));
}


// impl Solution {
//     pub fn minimum_time(n: i32, edges: Vec<Vec<i32>>, disappear: Vec<i32>) -> Vec<i32> {
//         let mut adj = vec![Vec::new(); n as usize];
//         for edge in edges.iter() {
//             let u = edge[0] as usize;
//             let v = edge[1] as usize;
//             let length = edge[2];
//             adj[u].push((v, length));
//             adj[v].push((u, length));
//         }
        
//         let mut pq = BinaryHeap::new();
//         pq.push(Reverse((0, 0)));
//         let mut answer = vec![-1; n as usize];
//         answer[0] = 0;
        
//         while let Some(Reverse((t, u))) = pq.pop() {
//             if t != answer[u] {
//                 continue;
//             }
//             for &(v, length) in adj[u].iter() {
//                 if t + length < disappear[v] && (answer[v] == -1 || t + length < answer[v]) {
//                     pq.push(Reverse((t + length, v)));
//                     answer[v] = t + length;
//                 }
//             }
//         }
        
//         answer
//     }
// }
