struct Solution {}

impl Solution {
    pub fn cal_points(operations: Vec<String>) -> i32 {
        let mut ans = 0;
        let mut l = Vec::new();
        for s in operations {

            if s == "D" {
                l.push(l[l.len()-1] * 2);
            } else if s == "C" {
                l.remove(l.len()-1);
            } else if s == "+" {
                l.push(l[l.len()-2] + l[l.len() - 1]);
            } else {
                let n = s.parse::<i32>().unwrap();
                l.push(n);
            }
        }

        for i in l {
            ans += i;
        }
        return ans as i32;
    }
}

fn main() {
    println!("{}", Solution::cal_points(vec!["5".to_string(),"-2".to_string(),"4".to_string(),"C".to_string(),"D".to_string(),"9".to_string(),"+".to_string(),"+".to_string()]))
}