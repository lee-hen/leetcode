use rand::seq::SliceRandom;
use rand::thread_rng;

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct Interval {
    pub start: i32,
    pub end: i32,
}

impl Interval {
    #[inline]
    pub fn new(start: i32, end: i32) -> Self {
        Interval { start, end }
    }
}

pub fn employee_free_time(schedule: Vec<Vec<Interval>>) -> Vec<Interval> {
    let mut ints = Vec::new();

    for intervals in schedule {
        for interval in intervals {
            ints.push(interval);
        }
    }

    ints.shuffle(&mut thread_rng());
    ints.sort_unstable_by(|a, b| a.start.cmp(&b.start));

    let mut res = Vec::new();
    let mut end = ints[0].end;

    for int in ints.iter().skip(1) {
        if int.start > end {
            res.push(Interval {
                start: end,
                end: int.start,
            })
        }
        end = i32::max(end, int.end);
    }
    res
}

#[cfg(test)]
mod tests {
    use crate::{employee_free_time, Interval};

    #[test]
    fn test_cases() {
        let schedule = vec![
            vec![Interval { start: 1, end: 2 }, Interval { start: 5, end: 6 }],
            vec![Interval { start: 1, end: 3 }],
            vec![Interval { start: 4, end: 10 }],
        ];
        let expect = vec![Interval { start: 3, end: 4 }];
        assert_eq!(expect, employee_free_time(schedule));

        let schedule = vec![
            vec![Interval { start: 1, end: 3 }, Interval { start: 6, end: 7 }],
            vec![Interval { start: 2, end: 4 }],
            vec![
                Interval { start: 2, end: 5 },
                Interval { start: 9, end: 12 },
            ],
        ];
        let expect = vec![Interval { start: 5, end: 6 }, Interval { start: 7, end: 9 }];
        assert_eq!(expect, employee_free_time(schedule));
    }
}
