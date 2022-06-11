pub fn number_of_patterns(m: i32, n: i32) -> i32 {
    let mut skips: Vec<Vec<usize>> = vec![Vec::new(); 10];
    #[allow(clippy::needless_range_loop)]
    for i in 0..10 {
        skips[i] = vec![0; 10];
    }

    (skips[1][3], skips[3][1]) = (2, 2);
    (skips[1][7], skips[7][1]) = (4, 4);
    (skips[3][9], skips[9][3]) = (6, 6);
    (skips[7][9], skips[9][7]) = (8, 8);
    (
        skips[1][9],
        skips[9][1],
        skips[2][8],
        skips[8][2],
        skips[3][7],
        skips[7][3],
        skips[4][6],
        skips[6][4],
    ) = (5, 5, 5, 5, 5, 5, 5, 5);

    fn dfs(visited: &mut Vec<bool>, skips: &Vec<Vec<usize>>, cur: usize, remain: i32) -> i32 {
        if remain < 0 {
            return 0;
        }
        if remain == 0 {
            return 1;
        }
        visited[cur] = true;
        let mut res = 0;
        for i in 1..=9 {
            if !visited[i] && (skips[cur][i] == 0 || visited[skips[cur][i]]) {
                res += dfs(visited, skips, i, remain - 1)
            }
        }
        visited[cur] = false;
        res
    }

    let mut visited: Vec<bool> = vec![false; 10];
    let mut res = 0;

    for i in m..=n {
        res += dfs(&mut visited, &skips, 1, i - 1) * 4; // 1, 3, 7, 9 are symmetric
        res += dfs(&mut visited, &skips, 2, i - 1) * 4; // 2, 4, 6, 8 are symmetric
        res += dfs(&mut visited, &skips, 5, i - 1); // 5
    }
    res
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case() {
        assert_eq!(number_of_patterns(1, 1), 9);
        assert_eq!(number_of_patterns(1, 2), 65);
        assert_eq!(number_of_patterns(2, 2), 56);
        assert_eq!(number_of_patterns(9, 9), 140704);
        assert_eq!(number_of_patterns(3, 3), 320);
        assert_eq!(number_of_patterns(1, 3), 385);
    }
}
