use std::cmp;

pub fn max_profit(prices: Vec<i32>) -> i32 {
    if prices.len() == 0 {
        return 0;
    }

    let mut max_so_far = i32::MIN;
    let mut max_profit = 0;

    for i in 1..prices.len() {
        max_so_far = cmp::max(max_so_far, -prices[i - 1]);
        max_profit = cmp::max(max_profit, max_so_far + prices[i]);
    }

    max_profit
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case() {
        let prices = vec![7, 1, 5, 3, 6, 4];
        assert_eq!(max_profit(prices), 5);

        let prices = vec![7, 6, 4, 3, 1];
        assert_eq!(max_profit(prices), 0);
    }
}
