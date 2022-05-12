use std::collections::HashMap;

pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut memorize: HashMap<i32, i32> = HashMap::new();

    for (i, y) in nums.into_iter().enumerate() {
        let x = target - y;

        if memorize.contains_key(&x) {
            let j = memorize.get(&x);

            let vec = match j {
                Some(j) => vec![i as i32, *j],
                None => vec![],
            };

            return vec;
        }
        memorize.insert(y, i as i32);
    }

    vec![]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case1() {
        let nums = vec![2, 7, 11, 15];
        let target = 9;
        let expected: Vec<i32> = vec![1, 0];

        assert_eq!(two_sum(nums, target), expected);
    }

    #[test]
    fn test_case2() {
        let nums = vec![3, 2, 4];
        let target = 6;
        let expected: Vec<i32> = vec![2, 1];

        assert_eq!(two_sum(nums, target), expected);
    }

    #[test]
    fn test_case3() {
        let nums = vec![3, 3];
        let target = 6;
        let expected: Vec<i32> = vec![1, 0];

        assert_eq!(two_sum(nums, target), expected);
    }
}
