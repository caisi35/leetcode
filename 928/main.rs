
struct Solution {}

impl Solution {
    fn dfs(graph: &Vec<Vec<i32>>, init_set: &Vec<i32>, infe_set: &mut Vec<i32>, v: usize) {
        let n = graph.len() as i32;
        for u in 0..n as usize {
            if graph[v][u] == 0 || init_set[u] == 1 || infe_set[u] == 1 {
                continue
            }
            infe_set[u] = 1;
            Self::dfs(graph, init_set, infe_set, u);
        }
    }

    pub fn min_malware_spread(graph: Vec<Vec<i32>>, initial: Vec<i32>) -> i32 {
        let n = graph.len();
        let mut init_set : Vec<i32> = vec![0;n];
        for &v in initial.iter() {
            init_set[v as usize] = 1;
        }
        let mut infe_by: Vec<Vec<i32>> = vec![Vec::new(); n];
        for &v in initial.iter() {
            let mut infe_set: Vec<i32> = vec![0;n];
            Self::dfs(&graph, & init_set, &mut infe_set, v as usize);
            for u in 0..n {
                if infe_set[u] == 1 {
                    infe_by[u].push(v);
                }
            }
        }
        let mut count: Vec<i32> = vec![0;n];
        for u in 0..n {
            if infe_by[u].len() == 1 {
                count[infe_by[u][0] as usize] += 1;
            }
        }
        let mut res = initial[0];
        for &v in initial.iter() {
            if count[ v as usize] > count[res as usize] || count[v as usize] == count[res as usize] && v < res {
                res = v;
            }
        }
        return res;
    }
}

fn main() {
    println!("main");
    let initial = vec![1, 2];
    let g = vec![vec![1, 1, 1], vec![1, 1, 1], vec![1, 1, 1]];
    println!("{}", Solution::min_malware_spread(g, initial));
}
