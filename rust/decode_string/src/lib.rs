use std::str::FromStr;

pub fn decode_string(s: String) -> String {
    fn decode_string(string: &String, idx: &mut usize) -> String {
        let mut result = String::new();
        let mut s = string.get(*idx..*idx + 1).unwrap();
        while *idx < string.len() && s != "]" {
            let c = char::from_str(s).unwrap();

            if !c.is_digit(10) {
                result.push_str(s);
                *idx += 1;
            } else {
                let mut k = 0;

                while *idx < string.len() {
                    let s = string.get(*idx..*idx + 1).unwrap();
                    let c = char::from_str(s).unwrap();
                    if !c.is_digit(10) {
                        break;
                    }

                    k = k * 10 + (s.as_bytes()[0] - 48);
                    *idx += 1;
                }

                *idx += 1;
                let decode_string = decode_string(string, idx);
                *idx += 1;

                while k > 0 {
                    result.push_str(&decode_string);
                    k -= 1;
                }
            }

            if let Some(_s) = string.get(*idx..*idx + 1) {
                s = _s;
            }
        }

        result
    }

    let mut idx: usize = 0;
    decode_string(&s, &mut idx)
}

#[cfg(test)]
mod tests {
    use crate::decode_string;

    #[test]
    fn it_works() {
        let s = String::from("3[a]2[bc]");
        let r = decode_string(s);
        assert_eq!(r, "aaabcbc");

        let s = String::from("3[a2[c]]");
        let r = decode_string(s);
        assert_eq!(r, "accaccacc");

        let s = String::from("2[abc]3[cd]ef");
        let r = decode_string(s);
        assert_eq!(r, "abcabccdcdcdef");

        let s = String::from("3[a]2[bc]");
        let r = decode_string(s);
        assert_eq!(r, "aaabcbc");

        let s = String::from("3[a2[c]]");
        let r = decode_string(s);
        assert_eq!(r, "accaccacc");

        let s = String::from("2[abc]3[cd]ef");
        let r = decode_string(s);
        assert_eq!(r, "abcabccdcdcdef");
    }
}
