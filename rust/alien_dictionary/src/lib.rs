use std::borrow::{Borrow, BorrowMut};
use std::cmp::min;
use std::collections::{HashMap, VecDeque};
use std::ops::Deref;

pub fn alien_order(words: Vec<String>) -> String {
    let mut adj_list = HashMap::new();
    let mut counts = HashMap::new();
    for (_, word) in words.iter().enumerate() {
        for (_, b) in word.as_bytes().iter().enumerate() {
            counts.insert(b, 0);
            adj_list.insert(b, Vec::new());
        }
    }

    for i in 0..words.len() - 1 {
        let word1 = words.get(i).unwrap();
        let word2 = words.get(i + 1).unwrap();

        if word1.len() > word2.len() && word1.as_bytes().starts_with(word2.as_bytes()) {
            return String::new();
        }

        for j in 0..min(word1.len(), word2.len()) {
            let b1 = word1.as_bytes().get(j).unwrap();
            let b2 = word2.as_bytes().get(j).unwrap();

            if b1 != b2 {
                adj_list.get_mut(b1).unwrap().push(b2);
                let val = *counts.get(b2).unwrap();
                counts.insert(b2, val + 1);

                break;
            }
        }
    }

    let mut queue = VecDeque::new();
    for (b, count) in counts.clone() {
        if count == 0 {
            queue.push_back(b);
        }
    }

    let mut out_put = Vec::new();
    while !queue.is_empty() {
        let front_byte = queue.pop_front().unwrap();
        out_put.push(*front_byte);

        for (_, n) in adj_list.get(front_byte).unwrap().iter().enumerate() {
            let val = *counts.get(n).unwrap();
            let val = val - 1;
            if val == 0 {
                queue.push_back(n);
            }
            counts.insert(n, val);
        }
    }

    if out_put.len() < counts.len() {
        String::new()
    } else {
        match String::from_utf8(out_put) {
            Ok(str) => str,
            Err(..) => String::new(),
        }
    }
}

#[derive(PartialEq, Debug)]
enum State {
    Visiting,
    Visited,
}

pub fn alien_order_dfs(words: Vec<String>) -> String {
    let graph = match build_graph(words.borrow()) {
        Ok(g) => g,
        Err(..) => return String::from(""),
    };

    let mut output = Vec::new();
    let mut mark = HashMap::new();
    for b in graph.keys() {
        let b = b.deref();
        let has_cycle = dfs(b, output.as_mut(), mark.borrow_mut(), graph.borrow());

        if has_cycle {
            return String::from("");
        }
    }

    match String::from_utf8(output) {
        Ok(str) => str,
        Err(..) => String::from(""),
    }
}

// fn build_graph<'a>(words: &'a Vec<String>) -> Result<HashMap<&'a u8, Vec<&'a u8>>, String> {
fn build_graph(words: &Vec<String>) -> Result<HashMap<&u8, Vec<&u8>>, String> {
    let mut reverse_adj_list = HashMap::new();
    for (_, word) in words.iter().enumerate() {
        for (_, b) in word.as_bytes().iter().enumerate() {
            reverse_adj_list.insert(b, Vec::new());
        }
    }

    for i in 1..words.len() {
        let word1 = words.get(i - 1).unwrap();
        let word2 = words.get(i).unwrap();

        if word1.len() > word2.len() && word1.as_bytes().starts_with(word2.as_bytes()) {
            return Err(String::new());
        }

        for j in 0..min(word1.len(), word2.len()) {
            let b1 = word1.as_bytes().get(j).unwrap();
            let b2 = word2.as_bytes().get(j).unwrap();

            if b1 != b2 {
                reverse_adj_list.get_mut(b2).unwrap().push(b1);
                break;
            }
        }
    }
    Ok(reverse_adj_list)
}

fn dfs(
    b: &u8,
    output: &mut Vec<u8>,
    mark: &mut HashMap<u8, State>,
    g: &HashMap<&u8, Vec<&u8>>,
) -> bool {
    if mark.contains_key(b) {
        return *mark.get(b).unwrap() == State::Visiting;
    }

    mark.insert(*b, State::Visiting);
    for (_, n) in g.get(b).unwrap().iter().enumerate() {
        let has_cycle = dfs(n.deref(), output, mark, g);
        if has_cycle {
            return true;
        }
    }

    output.push(*b);
    mark.insert(*b, State::Visited);
    false
}

#[cfg(test)]
mod tests {
    use crate::alien_order;
    use crate::alien_order_dfs;

    #[test]
    fn test_case1() {
        let words1 = vec![
            "wrt".to_string(),
            "wrf".to_string(),
            "er".to_string(),
            "ett".to_string(),
            "rftt".to_string(),
        ];
        let words2 = words1.clone();
        let expected = "wertf";
        assert_eq!(alien_order_dfs(words1), expected);
        assert_eq!(alien_order(words2), expected);
    }

    #[test]
    fn test_case2() {
        let words1 = vec![String::from("z"), String::from("x")];
        let words2 = words1.clone();
        assert_eq!(alien_order_dfs(words1), "zx");
        assert_eq!(alien_order(words2), "zx");
    }

    #[test]
    fn test_case3() {
        let words1 = vec![String::from("z"), String::from("x"), String::from("z")];
        let words2 = words1.clone();

        assert_eq!(alien_order_dfs(words1), "");
        assert_eq!(alien_order(words2), "");
    }

    #[test]
    fn test_case4() {
        let words = vec![String::from("abbb"), String::from("abc")];
        let actual = alien_order_dfs(words);

        match actual {
            a if a == *"bca" => assert_eq!(a, "bca"),
            a if a == *"abc" => assert_eq!(a, "abc"),
            _ => assert_eq!(actual, "bac"),
        }
    }

    #[test]
    fn test_case5() {
        let words1 = vec![String::from("z"), String::from("z")];
        let words2 = words1.clone();
        assert_eq!(alien_order_dfs(words1), "z");
        assert_eq!(alien_order(words2), "z");
    }
}
