struct Solution {}

impl Solution {
    pub fn count_good_nodes(edges: Vec<Vec<i32>>) -> i32 {
        let n = edges.len() + 1;
        let mut g = vec![vec![]; n];
        for edge in edges {
            let (x, y) = (edge[0] as usize, edge[1] as usize);
            g[x].push(y);
            g[y].push(x);
        }
        let mut res = 0;
        fn dfs(node: usize, parent: isize, g:&Vec<Vec<usize>>, res: &mut i32) -> usize {
            let mut valid = true;
            let mut treeSize = 0;
            let mut subTreeSize = 0;
            for &child in &g[node] {
                if child != parent as usize {
                    let size = dfs(child, node as isize, g, res);
                    if subTreeSize == 0 {
                        subTreeSize = size;
                    } else if size != subTreeSize {
                        valid = false;
                    }
                    treeSize += size;
                }
            }
            if valid {
                *res+=1;
            }
            return treeSize + 1;
        }
        dfs(0, -1, &g, &mut res);
        return res;
    }
}

fn main() {
    println!("{}", Solution::count_good_nodes(vec![
        vec![0,1],
        vec![0,2],
        vec![1,3],
        vec![1,4],
        vec![2,5],
        vec![2,6],
    ]))
}