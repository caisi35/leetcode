
struct Solution{}

impl Solution {
    pub fn find_missing_and_repeated_values(grid: Vec<Vec<i32>>) -> Vec<i32> {
        // println!("{}", grid.len()+1);
        let leng = grid.len() * grid.len() +1;
        // let len = grid.len() + 1;
        let mut l = vec![0; leng];
        // println!("{:#?}", l);
        for g in grid.iter() {
            for i in g.iter() {
                l[*i as usize] += 1;
            }
        }


        let mut ans: Vec<i32> = vec![0, 0];

        for (i, &v) in l.iter().enumerate() {
            if v == 2 {
                ans[0] = i as i32
            } else if v == 0 {
                ans[1] = i as i32
            }
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::find_missing_and_repeated_values(vec![
        vec![9,1,7],
        vec![8,9,2],
        vec![3,4,6]
    ]))
}