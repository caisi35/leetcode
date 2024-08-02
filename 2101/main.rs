use std::collections::{HashMap, VecDeque};
use std::cmp::max;

struct Solution {}

impl Solution {
    pub fn maximum_detonation(bombs: Vec<Vec<i32>>) -> i32 {
        let n = bombs.len();
        let is_conn = |u: usize, v: usize| -> bool {
            let dx = (bombs[u][0] - bombs[v][0]) as i64;
            let dy = (bombs[u][1] - bombs[v][1]) as i64;
            (bombs[u][2] as i64) * (bombs[u][2] as i64) >= (dx * dx + dy * dy)
        };
        let mut edges: HashMap<usize, Vec<usize>> = HashMap::new();
        for i in 0..n {
            for j in 0..n {
                if i != j && is_conn(i, j) {
                    edges.entry(i).or_insert(Vec::new()).push(j);
                }
            }
        }

        let mut res = 0;
        for i in 0..n {
            let mut visited = vec![0; n];
            let mut cnt = 1;
            let mut q = VecDeque::new();
            q.push_back(i);
            visited[i] = 1;
            while let Some(cidx) = q.pop_front() {
                for &nidx in edges.get(&cidx).unwrap_or(&vec![]) {
                    if visited[nidx] == 1 {
                        continue
                    }
                    cnt += 1;
                    q.push_back(nidx);
                    visited[nidx] = 1;
                }
            }
            res = max(res, cnt);
        }
        res
    }
}



// impl Solution {
//     pub fn maximum_detonation(bombs: Vec<Vec<i32>>) -> i32 {
//         let n = bombs.len();
//     // 判断炸弹 u 能否引爆炸弹 v
//         let is_connected = |u: usize, v: usize| -> bool {
//             let dx = (bombs[u][0] - bombs[v][0]) as i64;
//             let dy = (bombs[u][1] - bombs[v][1]) as i64;
//             (bombs[u][2] as i64) * (bombs[u][2] as i64) >= (dx * dx + dy * dy)
//         };

//         // 维护引爆关系有向图
//         let mut edges: HashMap<usize, Vec<usize>> = HashMap::new();
//         for i in 0..n {
//             for j in 0..n {
//                 if i != j && is_connected(i, j) {
//                     edges.entry(i).or_insert(Vec::new()).push(j);
//                 }
//             }
//         }

//         let mut res = 0; // 最多引爆数量
//         for i in 0..n {
//             // 遍历每个炸弹，广度优先搜索计算该炸弹可引爆的数量，并维护最大值
//             let mut visited = vec![0; n];
//             let mut cnt = 1;
//             let mut q = VecDeque::new();
//             q.push_back(i);
//             visited[i] = 1;
//             while let Some(cidx) = q.pop_front() {
//                 for &nidx in edges.get(&cidx).unwrap_or(&vec![]) {
//                     if visited[nidx] == 1 {
//                         continue;
//                     }
//                     cnt += 1;
//                     q.push_back(nidx);
//                     visited[nidx] = 1;
//                 }
//             }
//             res = max(res, cnt);
//         }
//         res
//     }
// }



fn main() {
    println!("{}", Solution::maximum_detonation(vec![
        vec![2,1,3],
        vec![6,1,4]
    ]))
}