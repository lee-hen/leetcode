use std::collections::HashMap;
use std::str;

pub fn crack_safe(n: i32, k: i32) -> String {
    let mut string = String::new();
    for _ in 0..n {
       string.push_str("0");
    }

    let mut visited = HashMap::new();
    visited.insert(string.clone(), true);
    let limit = i32::pow(k, n as u32);

    dfs(&mut string, k as u8, n, limit, &mut visited);
    string
}

fn dfs(
    string: &mut String,
    k: u8,
    n: i32,
    limit: i32,
    visited: &mut HashMap<String, bool>,
) -> bool {
    if visited.len() == limit as usize {
        return true;
    }

    let mut c: u8 = 48;
    let last = string
        .get((string.len() - (n as usize) + 1)..string.len())
        .unwrap()
        .to_string();
    while c < 48 + k {
        let mut new_string = last.clone();
        new_string.push_str(str::from_utf8(&[c]).unwrap());
        if !visited.contains_key(&new_string) {
            visited.insert(new_string.clone(), true);
            string.push_str(str::from_utf8(&[c]).unwrap());
            if dfs(string, k, n, limit, visited) {
                return true;
            }
            visited.remove(&new_string);
            string.pop();
        }
        c += 1;
    }

    false
}

#[cfg(test)]
mod tests {
    use crate::crack_safe;

    #[test]
    fn test_case() {
        let n = 2;
        let k = 2;
        let result = crack_safe(n, k);

        assert!(vec![
            String::from("01100"),
            String::from("10011"),
            String::from("11001"),
            String::from("00110"),
        ]
        .contains(&result));
    }
}
