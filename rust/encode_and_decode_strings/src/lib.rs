#![allow(dead_code)]
struct Codec {}

impl Codec {
    fn new() -> Self {
        Self {}
    }

    fn encode(&self, strs: Vec<String>) -> String {
        let mut string = String::new();

        for s in strs {
            let bytes = s.as_bytes();
            string.push_str(&*format!("{:?}", bytes));
            string.push(';');
        }
        string
    }

    fn decode(&self, s: String) -> Vec<String> {
        let s = s.split(';');
        let mut result = Vec::new();

        for c in s {
            if !c.is_empty() {
                let mut string = String::new();
                let byte_of_strings = c.get(1..c.len() - 1).unwrap_or("").to_string();
                let byte_of_strings = byte_of_strings.split(", ");

                for byte_str in byte_of_strings {
                    if !byte_str.is_empty() {
                        let byte_number = byte_str.parse::<u8>().unwrap_or(0);
                        string.push_str(std::str::from_utf8(&[byte_number]).unwrap_or(""));
                    }
                }

                if !string.is_empty() {
                    result.push(string);
                } else {
                    result.push(String::new());
                }
            }
        }

        result
    }
}

#[cfg(test)]
mod tests {
    use crate::Codec;

    #[test]
    fn test_case1() {
        let dummy_input = vec![String::from("Hello"), String::from("World")];
        let code_c = Codec::new();
        assert_eq!(
            code_c.encode(dummy_input),
            "[72, 101, 108, 108, 111];[87, 111, 114, 108, 100];"
        );
    }
    #[test]
    fn test_case2() {
        let code_c = Codec::new();
        assert_eq!(
            code_c.decode(String::from(
                "[72, 101, 108, 108, 111];[87, 111, 114, 108, 100];"
            )),
            vec![String::from("Hello"), String::from("World")]
        );
    }

    #[test]
    fn test_case3() {
        let code_c = Codec::new();
        assert_eq!(
            code_c.decode(code_c.encode(vec![String::new()])),
            vec![String::new()]
        );
    }
}
