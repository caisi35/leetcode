struct Solution{}

impl Solution {
    pub fn shortest_distance_after_queries(n: i32, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut roads: Vec<i32> = (1..=n).collect();
        let mut res: Vec<i32> = Vec::new();
        let mut dist = n-1;
        for query in &queries {
            let mut k = roads[query[0] as usize];
            roads[query[0] as usize] = query[1];
            while k != -1 && k < query[1] {
                let t = roads[k as usize];
                roads[k as usize] = -1;
                k = t;
                dist -= 1;
            }
            res.push(dist);
        }
        res
    }
}

fn main() {
    println!("{:#?}", Solution::shortest_distance_after_queries(5, vec![vec![2,4], vec![0,2], vec![0,4]]));
}