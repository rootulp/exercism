use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut available_words: HashMap<&str, usize> = HashMap::new();
    let mut needed_words: HashMap<&str, usize> = HashMap::new();
    for word in magazine {
        let count = available_words.entry(word).or_insert(0);
        *count += 1;
    }

    for word in note {
        let count = needed_words.entry(word).or_insert(0);
        *count += 1;
    }

    for (word, count) in needed_words {
        if *available_words.entry(word).or_default() < count {
            return false;
        }
    }

    true
}
