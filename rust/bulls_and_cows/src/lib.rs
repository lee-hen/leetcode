pub fn get_hint(secret: String, guess: String) -> String {
    let mut h = vec![0; 10];
    let n = guess.len();

    let secret = secret.as_bytes();
    let guess = guess.as_bytes();

    let mut bulls = 0;
    let mut cows = 0;

    for idx in 0..n {
        let s = secret[idx];
        let g = guess[idx];
        if s == g {
            bulls += 1;
        } else {
            let s = (s - 48) as usize;
            if h[s] < 0 {
                cows += 1
            }

            let g = (g - 48) as usize;
            if h[g] > 0 {
                cows += 1
            }

            h[s] += 1;
            h[g] -= 1;
        }
    }
    format!("{}A{}B", bulls, cows)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case() {
        let secret = String::from("1807");
        let guess = String::from("7810");
        assert_eq!(get_hint(secret, guess), "1A3B");
        let secret = String::from("1123");
        let guess = String::from("0111");
        assert_eq!(get_hint(secret, guess), "1A1B");
    }
}
