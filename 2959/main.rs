struct Solution {}

impl Solution {
    pub fn number_of_sets(n: i32, max_distance: i32, roads: Vec<Vec<i32>>) -> i32 {
        let n_usize = n as usize;
        let max_distance_usize = max_distance as usize;
        let mut res = 0;
        let mut opened = vec![0; n_usize];
        for mask in 0..(1 << n) {
            for i in 0..n_usize {
                opened[i] = mask & (1 << i);
            }
            let mut d = vec![vec![1000000; n_usize]; n_usize];
            for road in &roads {
                let i = road[0] as usize;
                let j = road[1] as usize;
                let r = road[2] as usize;
                if opened[i] > 0 && opened[j] > 0 {
                    d[i][j] = d[i][j].min(r);
                    d[j][i] = d[j][i].min(r);
                }
            }
            for k in 0..n_usize {
                if opened[k] > 0 {
                    for i in 0..n_usize {
                        if opened[i] > 0 {
                            for j in (i + 1)..n_usize {
                                if opened[j] > 0 {
                                    d[i][j] = d[i][j].min(d[i][k] + d[k][j]);
                                    d[j][i] = d[i][j];
                                }
                            }
                        }
                    }
                }
            }
            let mut good = true;
            for i in 0..n_usize {
                if opened[i] > 0 {
                    for j in (i + 1)..n_usize {
                        if opened[j] > 0 && d[i][j] > max_distance_usize {
                            good = false;
                            break;
                        }
                    }
                    if !good {
                        break;
                    }
                }
            }
            if good {
                res += 1;
            }
        }
        res
    }
}

fn main() {
    println!(
        "{}",
        Solution::number_of_sets(3, 5, vec![vec![0, 1, 2], vec![1, 2, 10], vec![0, 2, 10],])
    )
}
