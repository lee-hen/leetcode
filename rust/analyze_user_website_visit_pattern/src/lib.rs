use std::collections::HashMap;

pub fn most_visited_pattern(
    usernames: Vec<String>,
    timestamps: Vec<i32>,
    websites: Vec<String>,
) -> Vec<String> {
    let (usernames, websites) = get_sorted(usernames, &timestamps, websites);
    let timestamp_indices = get_timestamp_indices(&timestamps);

    let mut timestamps_by_name: HashMap<String, Vec<i32>> = HashMap::new();
    for (i, username) in usernames.iter().enumerate() {
        if timestamps_by_name.contains_key(username) {
            timestamps_by_name
                .get_mut(username)
                .unwrap()
                .push(timestamps[i]);
        } else {
            timestamps_by_name.insert(username.clone(), vec![timestamps[i]]);
        }
    }

    let timestamp_patterns = get_timestamp_patterns(timestamps_by_name);

    let mut patterns: HashMap<String, i32> = HashMap::new();
    let mut users: HashMap<String, String> = HashMap::new();

    for timestamp_pattern in timestamp_patterns {
        let (timestamp1, timestamp2, timestamp3) = (
            timestamp_pattern[0],
            timestamp_pattern[1],
            timestamp_pattern[2],
        );
        let (idx1, idx2, idx3) = (
            timestamp_indices.get(&timestamp1).unwrap(),
            timestamp_indices.get(&timestamp2).unwrap(),
            timestamp_indices.get(&timestamp3).unwrap(),
        );
        let pattern_string = format!(
            "{},{},{}",
            websites[idx1.clone()],
            websites[idx2.clone()],
            websites[idx3.clone()]
        );
        if users
            .get(pattern_string.as_str())
            .unwrap_or(&String::from(""))
            .clone()
            != usernames[idx1.clone()]
        {
            patterns.insert(
                pattern_string.clone(),
                patterns.get(pattern_string.as_str()).unwrap_or(&0) + 1,
            );
            users.insert(pattern_string, usernames[idx1.clone()].clone());
        }
    }

    let mut most_visited = -1;
    let mut pattern_string = String::new();
    let mut pattern: Vec<String> = Vec::new();

    for (pattern_str, score) in patterns {
        if score > most_visited
            || score == most_visited
                && !pattern_string.is_empty()
                && pattern_str.le(&pattern_string)
        {
            most_visited = score;
            pattern_string = pattern_str;
            pattern = pattern_string.split(",").map(|p| String::from(p)).collect();
        }
    }
    pattern
}

fn get_timestamp_patterns(timestamps_by_name: HashMap<String, Vec<i32>>) -> Vec<Vec<i32>> {
    // let mut patterns: Vec<Vec<i32>> = Vec::new();
    // for (_, timestamps) in timestamps_by_name {
    //     if timestamps.len() == 3 {
    //         patterns.push(timestamps);
    //     } else if timestamps.len() > 3 {
    //         let timestamp_patterns = generate_timestamp_patterns(&timestamps);
    //         for timestamp_pattern in timestamp_patterns {
    //             patterns.push(timestamp_pattern);
    //         }
    //     }
    // }

    let patterns = Vec::new();
    let patterns =
        timestamps_by_name
            .into_values()
            .into_iter()
            .fold(patterns, |mut p, timestamps| {
                if timestamps.len() == 3 {
                    p.push(timestamps);
                } else if timestamps.len() > 3 {
                    let timestamp_patterns = generate_timestamp_patterns(&timestamps);
                    for timestamp_pattern in timestamp_patterns {
                        p.push(timestamp_pattern);
                    }
                }
                p
            });
    patterns
}

fn generate_timestamp_patterns(timestamp: &Vec<i32>) -> Vec<Vec<i32>> {
    let mut patterns = Vec::new();

    for i in 0..timestamp.len() - 2 {
        for j in i + 1..timestamp.len() - 1 {
            for k in j + 1..timestamp.len() {
                let pattern = vec![timestamp[i], timestamp[j], timestamp[k]];
                patterns.push(pattern);
            }
        }
    }
    patterns
}

fn get_sorted(
    usernames: Vec<String>,
    timestamps: &Vec<i32>,
    websites: Vec<String>,
) -> (Vec<String>, Vec<String>) {
    let timestamps = timestamps;
    let timestamp_indices = get_timestamp_indices(&timestamps);
    let mut timestamps = timestamps.to_vec();
    timestamps.sort();

    let mut sorted_usernames: Vec<String> = vec![String::new(); usernames.len()];
    let mut sorted_websites: Vec<String> = vec![String::new(); websites.len()];

    for (i, timestamp) in timestamps.iter().enumerate() {
        sorted_usernames[i] = usernames[timestamp_indices.get(timestamp).unwrap().clone()].clone();
        sorted_websites[i] = websites[timestamp_indices.get(timestamp).unwrap().clone()].clone();
    }

    (sorted_usernames, sorted_websites)
}

fn get_timestamp_indices(timestamps: &Vec<i32>) -> HashMap<i32, usize> {
    let mut timestamp_indices = HashMap::new();
    for (i, timestamp) in timestamps.iter().enumerate() {
        timestamp_indices.insert(timestamp.clone(), i);
    }

    timestamp_indices
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_case1() {
        let username = vec![
            "joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary",
        ];
        let timestamp = vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
        let website = vec![
            "home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career",
        ];
        let username: Vec<String> = username.into_iter().map(|s| String::from(s)).collect();
        let website: Vec<String> = website.into_iter().map(|s| String::from(s)).collect();
        let expected: Vec<String> = vec!["home", "about", "career"]
            .into_iter()
            .map(|s| String::from(s))
            .collect();
        assert_eq!(most_visited_pattern(username, timestamp, website), expected);
    }

    #[test]
    fn test_case2() {
        let username = vec!["ua", "ua", "ua", "ub", "ub", "ub"];
        let timestamp = vec![1, 2, 3, 4, 5, 6];
        let website = vec!["a", "b", "c", "a", "b", "a"];
        let username: Vec<String> = username.into_iter().map(|s| String::from(s)).collect();
        let website: Vec<String> = website.into_iter().map(|s| String::from(s)).collect();
        let expected: Vec<String> = vec!["a", "b", "a"]
            .into_iter()
            .map(|s| String::from(s))
            .collect();
        assert_eq!(most_visited_pattern(username, timestamp, website), expected);
    }

    #[test]
    fn test_case3() {
        let username = vec!["dowg", "dowg", "dowg"];
        let timestamp = vec![158931262, 562600350, 148438945];
        let website = vec!["y", "loedo", "y"];
        let username: Vec<String> = username.into_iter().map(|s| String::from(s)).collect();
        let website: Vec<String> = website.into_iter().map(|s| String::from(s)).collect();
        let expected: Vec<String> = vec!["y", "y", "loedo"]
            .into_iter()
            .map(|s| String::from(s))
            .collect();
        assert_eq!(most_visited_pattern(username, timestamp, website), expected);
    }

    #[test]
    fn test_case4() {
        let username = vec![
            "h", "eiy", "cq", "h", "cq", "txldsscx", "cq", "txldsscx", "h", "cq", "cq",
        ];
        let timestamp = vec![
            527896567, 334462937, 517687281, 134127993, 859112386, 159548699, 51100299, 444082139,
            926837079, 317455832, 411747930,
        ];
        let website = vec![
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "hibympufi",
            "yljmntrclw",
            "hibympufi",
            "yljmntrclw",
        ];
        let username: Vec<String> = username.into_iter().map(|s| String::from(s)).collect();
        let website: Vec<String> = website.into_iter().map(|s| String::from(s)).collect();
        let expected: Vec<String> = vec!["hibympufi", "hibympufi", "yljmntrclw"]
            .into_iter()
            .map(|s| String::from(s))
            .collect();
        assert_eq!(most_visited_pattern(username, timestamp, website), expected);
    }

    #[test]
    fn test_case5() {
        let username = vec![
            "pdtkrjhd", "pdtkrjhd", "pdtkrjhd", "pdtkrjhd", "pdtkrjhd", "pdtkrjhd",
        ];
        let timestamp = vec![
            210984153, 262799291, 958396687, 605779010, 373702273, 205190519,
        ];
        let website = vec!["xgriygexlk", "qs", "rugydl", "bkrok", "canlv", "cahgsobjjs"];
        let username: Vec<String> = username.into_iter().map(|s| String::from(s)).collect();
        let website: Vec<String> = website.into_iter().map(|s| String::from(s)).collect();
        let expected: Vec<String> = vec!["cahgsobjjs", "bkrok", "rugydl"]
            .into_iter()
            .map(|s| String::from(s))
            .collect();
        assert_eq!(most_visited_pattern(username, timestamp, website), expected);
    }

    #[test]
    fn test_case6() {
        let username = vec![
            "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c", "c",
        ];
        let timestamp = vec![
            192769792, 207063377, 333805625, 890700372, 64199933, 245678250, 69530300, 833864121,
            527969074, 492534187, 49153037, 143138463, 163274379,
        ];
        let website = vec![
            "jsips",
            "zkamv",
            "osajva",
            "bxbldeqhz",
            "zkamv",
            "osajva",
            "zkamv",
            "osajva",
            "zkamv",
            "zkamv",
            "zkamv",
            "zkamv",
            "jsips",
        ];
        let username: Vec<String> = username.into_iter().map(|s| String::from(s)).collect();
        let website: Vec<String> = website.into_iter().map(|s| String::from(s)).collect();
        let expected: Vec<String> = vec!["jsips", "jsips", "bxbldeqhz"]
            .into_iter()
            .map(|s| String::from(s))
            .collect();
        assert_eq!(most_visited_pattern(username, timestamp, website), expected);
    }

    #[test]
    fn test_generate_timestamp_patterns() {
        let expected = vec![
            vec![4, 5, 6],
            vec![4, 5, 7],
            vec![4, 5, 8],
            vec![4, 6, 7],
            vec![4, 6, 8],
            vec![4, 7, 8],
            vec![5, 6, 7],
            vec![5, 6, 8],
            vec![5, 7, 8],
            vec![6, 7, 8],
        ];
        assert_eq!(generate_timestamp_patterns(&vec![4, 5, 6, 7, 8]), expected);
    }
}
