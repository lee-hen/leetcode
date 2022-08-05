#![allow(dead_code)]

use std::cell::RefCell;
use std::collections::{HashMap, HashSet};
use std::rc::Rc;

struct AutocompleteSystem {
    string_to_search: String,
    prefix_db: HashMap<String, PrefixDB>,
}

struct PrefixDB {
    sentences: Vec<String>,
    hot_sentences: Vec<Rc<RefCell<HotSentence>>>,
    hot_sentence: Option<Rc<RefCell<HotSentence>>>,
}

impl PrefixDB {
    fn new() -> Self {
        Self {
            sentences: Vec::new(),
            hot_sentences: Vec::new(),
            hot_sentence: None,
        }
    }
}

#[derive(Hash, Eq, PartialEq, Debug)]
struct HotSentence {
    word: String,
    hot: i32,
}

impl HotSentence {
    fn new(word: String, hot: i32) -> Self {
        Self { word, hot }
    }
}

impl AutocompleteSystem {
    fn new(sentences: Vec<String>, times: Vec<i32>) -> Self {
        let mut hot_sentences = build_hot_sentences(sentences, times);
        sort_hot_sentences(&mut hot_sentences);
        Self {
            string_to_search: String::new(),
            prefix_db: build_prefix_db(&hot_sentences),
        }
    }

    fn input(&mut self, c: char) -> Vec<String> {
        if c == '#' {
            self.update_prefix_db(self.string_to_search.clone());
            self.string_to_search.clear();
            return Vec::new();
        }

        self.string_to_search.push_str(&String::from(c));
        if self.prefix_db.contains_key(&self.string_to_search) {
            return self
                .prefix_db
                .get(&self.string_to_search)
                .unwrap()
                .sentences
                .clone();
        }

        Vec::new()
    }

    fn update_prefix_db(&mut self, word: String) {
        let mut new_hot_sentence = None;
        if self.prefix_db.contains_key(&word) {
            if self.prefix_db.get(&word).unwrap().hot_sentence == None {
                new_hot_sentence = Some(Rc::new(RefCell::new(HotSentence::new(word.clone(), 1))));
            } else {
                let hot_sentence = match &self.prefix_db.get_mut(&word).unwrap().hot_sentence {
                    Some(s) => s,
                    None => return,
                };

                hot_sentence.as_ref().borrow_mut().hot += &1;
            }
        } else {
            new_hot_sentence = Some(Rc::new(RefCell::new(HotSentence::new(word.clone(), 1))));
        }

        for j in 1..word.len() + 1 {
            let prefix = word.get(0..j).unwrap_or("").to_string();
            if !self.prefix_db.contains_key(&prefix) {
                let mut db = PrefixDB::new();

                if new_hot_sentence.is_some() {
                    let new_hot_sentence = match &new_hot_sentence {
                        Some(s) => s,
                        None => return,
                    };

                    db.hot_sentences.push(Rc::clone(new_hot_sentence));
                }
                db.sentences.push(word.clone());
                self.prefix_db.insert(prefix.clone(), db);
            } else {
                if new_hot_sentence.is_some() {
                    let new_hot_sentence = match &new_hot_sentence {
                        Some(s) => s,
                        None => return,
                    };

                    self.prefix_db
                        .get_mut(&prefix)
                        .unwrap()
                        .hot_sentences
                        .push(Rc::clone(new_hot_sentence));
                }
                sort_hot_sentences(&mut self.prefix_db.get_mut(&prefix).unwrap().hot_sentences);

                let mut sentences = Vec::new();
                let mut i = 0;
                while i < 3 && i < self.prefix_db.get(&prefix).unwrap().hot_sentences.len() {
                    sentences.push(
                        self.prefix_db.get(&prefix).unwrap().hot_sentences[i]
                            .borrow()
                            .word
                            .clone(),
                    );
                    i += 1;
                }
                self.prefix_db.get_mut(&prefix).unwrap().sentences = sentences;
            }
        }
        if new_hot_sentence.is_some() {
            self.prefix_db.get_mut(&word).unwrap().hot_sentence = new_hot_sentence;
        }
    }
}

fn build_hot_sentences(sentences: Vec<String>, times: Vec<i32>) -> Vec<Rc<RefCell<HotSentence>>> {
    let mut hot_sentences = Vec::new();
    for i in 0..sentences.len() {
        hot_sentences.push(Rc::new(RefCell::new(HotSentence::new(
            sentences[i].clone(),
            times[i],
        ))));
    }

    hot_sentences
}

fn sort_hot_sentences(hot_sentences: &mut [Rc<RefCell<HotSentence>>]) {
    hot_sentences.sort_by(|a, b| {
        if a.borrow().hot == b.borrow().hot {
            return a.borrow().word.cmp(&b.borrow().word);
        }

        b.borrow().hot.cmp(&a.borrow().hot)
    });
}

fn build_prefix_db(hot_sentences: &[Rc<RefCell<HotSentence>>]) -> HashMap<String, PrefixDB> {
    let mut prefix_db = HashMap::new();
    let mut seen: HashMap<String, HashSet<HotSentence>> = HashMap::new();

    for (i, hot_sentence) in hot_sentences.iter().enumerate() {
        let mut j = 1;
        while j < hot_sentence.borrow().word.len() + 1 {
            let prefix = hot_sentence
                .borrow()
                .word
                .get(0..j)
                .unwrap_or("")
                .to_string();

            if !seen.contains_key(&prefix) {
                seen.insert(prefix.clone(), HashSet::new());
            }

            let hot_sentence = HotSentence::new(
                hot_sentence.borrow().word.clone(),
                hot_sentence.borrow().hot,
            );

            if !seen.get(&prefix).unwrap().contains(&hot_sentence) {
                seen.get_mut(&prefix).unwrap().insert(hot_sentence);
                if !prefix_db.contains_key(&prefix) {
                    prefix_db.insert(prefix.clone(), PrefixDB::new());
                }

                prefix_db
                    .get_mut(&prefix)
                    .unwrap()
                    .hot_sentences
                    .push(Rc::clone(&hot_sentences[i]));

                if prefix_db.get(&prefix).unwrap().sentences.len() < 3 {
                    prefix_db
                        .get_mut(&prefix)
                        .unwrap()
                        .sentences
                        .push(hot_sentences[i].borrow().word.clone());
                }
            }
            j += 1
        }
        prefix_db
            .get_mut(&hot_sentence.borrow().word)
            .unwrap()
            .hot_sentence = Some(hot_sentence.to_owned());
    }
    prefix_db
}

#[cfg(test)]
mod tests {
    use crate::AutocompleteSystem;

    #[test]
    fn test_case() {
        let mut s = AutocompleteSystem::new(
            vec![
                String::from("i love you"),
                String::from("island"),
                String::from("iroman"),
                String::from("i love leetcode"),
            ],
            vec![5, 3, 2, 2],
        );

        assert_eq!(
            vec![
                String::from("i love you"),
                String::from("island"),
                String::from("i love leetcode"),
            ],
            s.input('i')
        );
        assert_eq!(
            vec![String::from("i love you"), String::from("i love leetcode")],
            s.input(' ')
        );

        let expect: Vec<String> = Vec::new();
        assert_eq!(expect, s.input('a'));
        let expect: Vec<String> = Vec::new();
        assert_eq!(expect, s.input('#'));

        assert_eq!(
            vec![
                String::from("i love you"),
                String::from("island"),
                String::from("i love leetcode"),
            ],
            s.input('i')
        );

        assert_eq!(
            vec![
                String::from("i love you"),
                String::from("i love leetcode"),
                String::from("i a"),
            ],
            s.input(' ')
        );

        assert_eq!(vec![String::from("i a")], s.input('a'));

        let expect: Vec<String> = Vec::new();
        assert_eq!(expect, s.input('#'));

        assert_eq!(
            vec![
                String::from("i love you"),
                String::from("island"),
                String::from("i a"),
            ],
            s.input('i')
        );

        assert_eq!(
            vec![
                String::from("i love you"),
                String::from("i a"),
                String::from("i love leetcode"),
            ],
            s.input(' ')
        );

        assert_eq!(vec![String::from("i a")], s.input('a'));

        let expect: Vec<String> = Vec::new();
        assert_eq!(expect, s.input('#'));
    }
}
