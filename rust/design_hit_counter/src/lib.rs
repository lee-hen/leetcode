#![allow(dead_code)]
use std::collections::HashMap;

const FIVE_MINUTES: i32 = 300;

struct HitCounter {
    hits: HashMap<i32, i32>,

    last_hit_at: i32,
    five_minute_split_number: i32,
}

impl HitCounter {
    fn new() -> Self {
        Self {
            hits: HashMap::new(),
            last_hit_at: 0,
            five_minute_split_number: 0,
        }
    }

    fn hit(&mut self, timestamp: i32) {
        if timestamp / FIVE_MINUTES > self.five_minute_split_number + 1 {
            self.hits.clear();
        }

        self.hits.insert(
            timestamp,
            self.hits.get(&self.last_hit_at).unwrap_or(&0) + 1,
        );

        self.last_hit_at = timestamp;
        self.five_minute_split_number = timestamp / FIVE_MINUTES;
    }

    fn get_hits(&self, timestamp: i32) -> i32 {
        if timestamp < self.last_hit_at {
            panic!("not allowed");
        }

        if timestamp / FIVE_MINUTES > self.five_minute_split_number + 1 {
            return 0;
        }

        let mut target = 0;

        for time in self.hits.keys() {
            if *time <= timestamp - FIVE_MINUTES {
                target = i32::max(time.to_owned(), target);
            }
        }

        self.hits.get(&self.last_hit_at).unwrap_or(&0) - self.hits.get(&target).unwrap_or(&0)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case1() {
        let mut obj = HitCounter::new();
        obj.hit(1);
        obj.hit(2);
        obj.hit(3);

        assert_eq!(obj.get_hits(4), 3);

        obj.hit(300);
        assert_eq!(obj.get_hits(300), 4);
        assert_eq!(obj.get_hits(301), 3);
    }

    #[test]
    fn test_case2() {
        let mut obj = HitCounter::new();
        obj.hit(1);
        obj.hit(1);
        obj.hit(1);
        obj.hit(300);
        assert_eq!(obj.get_hits(300), 4);

        obj.hit(300);
        assert_eq!(obj.get_hits(300), 5);

        obj.hit(301);
        assert_eq!(obj.get_hits(301), 3);
    }

    #[test]
    fn test_case3() {
        let mut obj = HitCounter::new();
        obj.hit(100);
        obj.hit(282);
        obj.hit(411);
        obj.hit(609);
        obj.hit(620);
        obj.hit(744);
        assert_eq!(obj.get_hits(879), 3);

        obj.hit(956);
        assert_eq!(obj.get_hits(976), 2);

        obj.hit(998);
        assert_eq!(obj.get_hits(1055), 2);
    }
}
