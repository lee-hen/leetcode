use std::cmp;

pub fn max_area(height: Vec<i32>) -> i32 {
    if height.len() == 0 {
        return 0;
    }
    let mut left_idx = 0;
    let mut right_idx = height.len()-1;
    let mut surface_area = 0;

    while left_idx < right_idx {
        if height[left_idx] < height[right_idx] {
            surface_area = cmp::max((right_idx-left_idx) as i32 * height[left_idx], surface_area);
            left_idx += 1;
        } else {
            surface_area = cmp::max((right_idx-left_idx) as i32 * height[right_idx], surface_area);
            right_idx -= 1;
        }
    }
    surface_area
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cases() {
        let height = vec![1,8,6,2,5,4,8,3,7];
        assert_eq!(max_area(height), 49);

        let height = vec![1, 1];
        assert_eq!(max_area(height), 1);

        let height = vec![1, 1, 1, 1, 2, 2, 1, 1, 1];
        assert_eq!(max_area(height), 8);
    }
}
