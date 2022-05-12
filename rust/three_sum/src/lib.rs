use std::collections::HashMap;

#[derive(PartialEq, Eq, Hash, Debug)]
struct Triplet {
    num1: i32,
    num2: i32,
    num3: i32,
}

pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut nums = nums;
    nums.sort();
    let mut triplets: Vec<Vec<i32>> = Vec::new();
    let mut seen = HashMap::new();
    let length = nums.len();

    for (i, _) in nums.iter().enumerate() {
        let mut left = i + 1;
        let mut right = length - 1;
        while left < right {
            let current_sum = nums[i] + nums[left] + nums[right];
            if current_sum == 0 {
                let t = Triplet {
                    num1: nums[i],
                    num2: nums[left],
                    num3: nums[right],
                };
                let saw = seen.get(&t).unwrap_or(&false);
                if saw == &false {
                    triplets.push(vec![t.num1, t.num2, t.num3]);
                    seen.insert(t, true);
                }
                left += 1;
                right -= 1;
            } else if current_sum > 0 {
                right -= 1;
            } else {
                left += 1;
            }
        }
    }

    triplets
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case() {
        let expected = vec![vec![-1, -1, 2], vec![-1, 0, 1]];
        let nums = vec![-1, 0, 1, 2, -1, -4];
        assert_eq!(expected, three_sum(nums));
    }
}
