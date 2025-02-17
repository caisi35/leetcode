struct Solution {}

impl Solution {
    pub fn find_ball(grid: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ans = vec![0; grid[0].len()];
        for i in 0..ans.len() {
            let mut cur_col = i as i32;
            for j in grid.iter() {
                // println!("{:#?}", j);
                let cur = j[cur_col as usize];
                cur_col += cur;
                if cur_col < 0 || cur_col == ans.len() as i32 || j[cur_col as usize] != cur {
                    cur_col = -1;
                    break;
                }
            }
            ans[i] = cur_col;
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::find_ball(vec![vec![1,1,1,1,1,1],vec![-1,-1,-1,-1,-1,-1],vec![1,1,1,1,1,1],vec![-1,-1,-1,-1,-1,-1]]));
}