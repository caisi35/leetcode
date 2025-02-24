struct OrderedStream {
    data: Vec<Option<String>>,
    ptr: usize,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl OrderedStream {

    fn new(n: i32) -> Self {
        return OrderedStream{data: vec![None;1 + n as usize], ptr: 1}
    }
    
    fn insert(&mut self, id_key: i32, value: String) -> Vec<String> {
        let id = id_key as usize;
        self.data[id] = Some(value);
        let mut result = Vec::new();
        while self.ptr < self.data.len() {
            if let Some(val) = self.data[self.ptr].take() {
                result.push(val);
                self.ptr += 1;
            } else {
                break;
            }
        }
        result
    }
}

fn main() {
    let mut o = OrderedStream::new(5);
    println!("{:#?}", o.insert(3, "c".to_string()));
    println!("{:#?}", o.insert(1, "a".to_string()));

}