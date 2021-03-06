const SHARP: u8 = "#".as_bytes()[0];

pub fn backspace_compare(s: String, t: String) -> bool {
    let backspacing = |s: String| -> String {
        let mut buffer: Vec<u8> = Vec::new();
        for c in s.as_bytes() {
            if *c != SHARP {
                buffer.push(*c);
            } else if !buffer.is_empty() {
                buffer.pop();
            }
        }

        String::from_utf8(buffer).unwrap_or_default()
    };
    backspacing(s) == backspacing(t)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_case() {
        let s1 = String::from("ab#c");
        let s2 = String::from("ad#c");
        assert!(backspace_compare(s1, s2));

        let s1 = String::from("ab##");
        let s2 = String::from("c#d#");
        assert!(backspace_compare(s1, s2));

        let s1 = String::from("a#c");
        let s2 = String::from("b");
        assert!(!backspace_compare(s1, s2));
    }
}
