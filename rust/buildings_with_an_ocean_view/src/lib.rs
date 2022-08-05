pub fn find_buildings(heights: Vec<i32>) -> Vec<i32> {
    if heights.len() == 1 {
        return vec![0];
    }
    let mut ocean_view = vec![(heights.len() - 1) as i32];
    let mut max = heights[heights.len() - 1];
    for i in (0..=heights.len() - 2).rev() {
        if heights[i] > max {
            max = heights[i];
            ocean_view.push(i as i32);
        }
    }
    ocean_view.into_iter().rev().collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_cases() {
        assert_eq!(find_buildings(vec![4, 2, 3, 1]), vec![0, 2, 3]);
        assert_eq!(find_buildings(vec![4, 3, 2, 1]), vec![0, 1, 2, 3]);
        assert_eq!(find_buildings(vec![1, 3, 2, 4]), vec![3]);
        assert_eq!(find_buildings(vec![2, 2, 2, 2]), vec![3]);
        assert_eq!(find_buildings(vec![44]), vec![0]);
    }
}
