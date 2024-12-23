struct Solution {}

impl Solution {
    pub fn can_reach_corner(x_corner: i32, y_corner: i32, circles: Vec<Vec<i32>>) -> bool {
        fn point_in_circle(px: i32, py: i32, x: i32, y: i32, r: i32) -> bool {
            let dx = x as i128 - px as i128;
            let dy = y as i128 - py as i128;
            dx * dx + dy * dy <= (r as i128) * (r as i128)
        }

        fn circle_intersects_top_left(x: i32, y: i32, r: i32, x_corner: i32, y_corner: i32) -> bool {
            (x.abs() <= r && y >= 0 && y <= y_corner)
                || (x >= 0 && x <= x_corner && (y - y_corner).abs() <= r)
                || point_in_circle(0, y_corner, x, y, r)
        }

        fn circle_intersects_bottom_right(x: i32, y: i32, r: i32, x_corner: i32, y_corner: i32) -> bool {
            (y.abs() <= r && x >= 0 && x <= x_corner)
                || (y >= 0 && y <= y_corner && (x - x_corner).abs() <= r)
                || point_in_circle(x_corner, 0, x, y, r)
        }

        fn circles_intersect_in_rectangle(
            x1: i32, y1: i32, r1: i32,
            x2: i32, y2: i32, r2: i32,
            x_corner: i32, y_corner: i32,
        ) -> bool {
            let dx = x1 as i128 - x2 as i128;
            let dy = y1 as i128 - y2 as i128;
            let r_sum = r1 as i128 + r2 as i128;
            dx * dx + dy * dy <= r_sum * r_sum
                && (x1 as i128 * r2 as i128 + x2 as i128 * r1 as i128) < r_sum * x_corner as i128
                && (y1 as i128 * r2 as i128 + y2 as i128 * r1 as i128) < r_sum * y_corner as i128
        }

        fn dfs(i: usize, circles: &Vec<Vec<i32>>, x_corner: i32, y_corner: i32, visited: &mut [bool]) -> bool {
            let x1 = circles[i][0];
            let y1 = circles[i][1];
            let r1 = circles[i][2];
            if circle_intersects_bottom_right(x1, y1, r1, x_corner, y_corner) {
                return true;
            }
            visited[i] = true;
            for (j, vec) in circles.iter().enumerate() {
                let x2 = vec[0];
                let y2 = vec[1];
                let r2 = vec[2];
                if !visited[j]
                    && circles_intersect_in_rectangle(x1, y1, r1, x2, y2, r2, x_corner, y_corner)
                    && dfs(j, circles, x_corner, y_corner, visited) {
                    return true;
                }
            }
            false
        }

        let mut visited = vec![false; circles.len()];

        for (i, vec) in circles.iter().enumerate() {
            let x = vec[0];
            let y = vec[1];
            let r = vec[2];
            if point_in_circle(0, 0, x, y, r) || point_in_circle(x_corner, y_corner, x, y, r) {
                return false;
            }
            if !visited[i]
                && circle_intersects_top_left(x, y, r, x_corner, y_corner)
                && dfs(i, &circles, x_corner, y_corner, &mut visited) {
                return false;
            }
        }
        true
    }
}

fn main() {
    println!("{}", Solution::can_reach_corner(3,3, vec![vec![3,1,1]]))
}