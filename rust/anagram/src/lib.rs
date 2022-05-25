use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let mut anagrams: HashSet<&'a str> = HashSet::new();

    for possible in possible_anagrams {
        if is_anagram(word, possible) {
            anagrams.insert(possible);
        }
    }
    anagrams
}

fn is_anagram(word: &str, possible: &str) -> bool {
    if word == possible {
        return false;
    }

    return sorted_lowercase_chars(word) == sorted_lowercase_chars(possible);
}

fn sorted_lowercase_chars(word: &str) -> Vec<char> {
    let mut word_chars = word.to_lowercase().chars().collect::<Vec<char>>();
    word_chars.sort_unstable();

    return word_chars;
}
