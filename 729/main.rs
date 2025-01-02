
struct MyCalendar {
    l:  Vec<Vec<i32>>,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyCalendar {

    fn new() -> Self {
        MyCalendar{l: vec![]}
    }
    
    fn book(&mut self, start_time: i32, end_time: i32) -> bool {
        for i in &self.l {
            if start_time < i[1] && end_time > i[0] {
                return false;
            }
        }
        let _ = &self.l.push(vec![start_time, end_time]);
        true
    }
}


fn main () {
    let mut obj = MyCalendar::new();
    let ret_1: bool = obj.book(10, 20);
    let ret_2: bool = obj.book(15, 20);
    let ret_3: bool = obj.book(30, 50);
    println!("{}", ret_1);
    println!("{}", ret_2);
    println!("{}", ret_3);
}