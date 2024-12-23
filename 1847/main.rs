use std::collections::BTreeSet;

struct Solution {}

#[derive(Debug)]
struct Event {
    event_type: i32,
    size: i32,
    id: i32,
    origin: usize,
}

impl Event {
    fn new(event_type: i32, size: i32, id: i32, origin: usize) -> Self {
        Event {
            event_type,
            size,
            id,
            origin,
        }
    }
}

impl PartialEq for Event {
    fn eq(&self, other: &Self) -> bool {
        self.size == other.size && self.event_type == other.event_type 
    }
}

impl Eq for Event {}

impl PartialOrd for Event {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Event {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other.size.cmp(&self.size).then(self.event_type.cmp(&other.event_type))
    }
}

impl Solution {
    pub fn closest_room(rooms: Vec<Vec<i32>>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut events = Vec::new();
        for (i, room) in rooms.iter().enumerate() {
            events.push(Event::new(0, room[1], room[0], i));
        }
        for (i, query) in queries.iter().enumerate() {
            events.push(Event::new(1, query[1], query[0], i));
        }
        events.sort();
        let mut valid_rooms = BTreeSet::new();
        let mut ans = vec![-1; queries.len()];
        for event in events {
            if event.event_type == 0 {
                valid_rooms.insert(event.id);
            } else {
                let mut dist = i32::MAX;
                if let Some(&ceiling) = valid_rooms.range(event.id..).next() {
                    if ceiling - event.id < dist {
                        dist = ceiling - event.id;
                        ans[event.origin] = ceiling
                    }
                }
                if let Some(&floor) = valid_rooms.range(..event.id).next_back() {
                    if event.id - floor <= dist {
                        ans[event.origin] = floor;
                    }
                }
            }
        }
        ans
    }
}

fn main() {
    println!("{:#?}", Solution::closest_room(vec![vec![2,2], vec![1,2], vec![3,2]], 
        vec![vec![3,1], vec![3,3], vec![5,2]]));
}