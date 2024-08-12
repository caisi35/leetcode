


struct MagicDictionary {
    l: Vec<String>
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MagicDictionary {

    fn new() -> Self {
        return MagicDictionary{l: vec![]}
    }
    
    fn build_dict(&mut self, dictionary: Vec<String>) {
        self.l = dictionary
    }
    
    fn search(&self, search_word: String) -> bool {
        for w in self.l.iter() {
            if w.len() != search_word.len() {
                continue
            }
            let mut cnt = 0;
            for (i, ch) in w.chars().enumerate() {
                let cs = search_word.chars();
                for (j, c) in cs.enumerate() {
                    if i == j && ch != c {
                        cnt += 1;
                    }
                }
            }
            if cnt == 1 {
                return true
            }
        }
        false
    }
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * let obj = MagicDictionary::new();
 * obj.build_dict(dictionary);
 * let ret_2: bool = obj.search(searchWord);
 */
fn main() {
    let mut obj = MagicDictionary::new();
    obj.build_dict(vec!["hello".to_string(),"hallo".to_string(),"leetcode".to_string()]);
    let ret_2: bool = obj.search("hhllo".to_string());
    println!("{}", ret_2)
}