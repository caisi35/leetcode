
struct Solution{}

impl Solution {
    pub fn discount_prices(sentence: String, discount: i32) -> String {
        let words: Vec<&str> = sentence.split_whitespace().collect();
        let mut replace: Vec<String> = Vec::new();
        for word in words.iter() {
            if word.starts_with("$") && word.len() > 1 && Self::is_numeric(&word[1..]) {
                let price : f64 = word[1..].parse().unwrap();
                let discounted_price = price * (1.0 - discount as f64 / 100.0);
                replace.push(format!("${:.2}", discounted_price));
            } else {
                replace.push(word.to_string());
            }
        }
        replace.join(" ")
    }

    fn is_numeric(s: &str) -> bool {
        s.chars().all(|c| c.is_digit(10))
    }
}

fn main() {
    println!("{}", Solution::discount_prices("there are $1 $2 and 5$ candies in the shop".to_string(), 50))
}