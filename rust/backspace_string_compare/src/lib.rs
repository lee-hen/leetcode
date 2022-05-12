pub fn backspace_compare(s: String, t: String) -> bool {
    let backspacing = |s: String| -> String {
        let mut buffer: Vec<u8> = Vec::new();
        for c in s.as_bytes() {
            if *c != "#".as_bytes()[0] {
                buffer.push(c.clone());
            } else if buffer.len() > 0 {
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
        assert_eq!(backspace_compare(s1, s2), true);

        let s1 = String::from("ab##");
        let s2 = String::from("c#d#");
        assert_eq!(backspace_compare(s1, s2), true);

        let s1 = String::from("a#c");
        let s2 = String::from("b");
        assert_eq!(backspace_compare(s1, s2), false);
    }
}
