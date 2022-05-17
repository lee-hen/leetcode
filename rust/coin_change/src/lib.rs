use std::cmp;

pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
    let mut num_of_coins = vec![i32::MAX; (amount + 1) as usize];
    num_of_coins[0] = 0;

    for coin in coins {
        for a in 1..=amount {
            if a == coin {
                num_of_coins[a as usize] = 1;
            } else if coin < a {
                if num_of_coins[(a - coin) as usize] != i32::MAX {
                    num_of_coins[a as usize] = cmp::min(
                        num_of_coins[a as usize],
                        num_of_coins[(a - coin) as usize] + 1,
                    );
                }
            }
        }
    }

    if num_of_coins[num_of_coins.len() - 1] == i32::MAX {
        return -1;
    }

    num_of_coins[num_of_coins.len() - 1]
}

#[cfg(test)]
mod tests {
    use crate::coin_change;

    #[test]
    fn test_cases() {
        let coins = vec![1, 2, 5];
        let amount = 11;
        assert_eq!(coin_change(coins, amount), 3);

        let coins = vec![2];
        let amount = 3;
        assert_eq!(coin_change(coins, amount), -1);
    }
}
