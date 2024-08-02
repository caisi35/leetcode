struct Solution{}

impl Solution {
    pub fn min_rectangles_to_cover_points(mut points: Vec<Vec<i32>>, w: i32) -> i32 {
        let mut bound = -1;
        let mut ans = 0;
        points.sort();
        for point in points {
            if point[0] > bound {
                bound = point[0] + w;
                ans += 1;
            }
        }
        ans
    }
}

fn main() {
    println!("{}", Solution::min_rectangles_to_cover_points(vec![
        vec![2,3],
        vec![1, 2],
    ], 0))
}