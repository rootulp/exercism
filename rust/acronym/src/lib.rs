pub fn abbreviate(phrase: &str) -> String {
    if phrase.is_empty() {
        return String::new();
    }
    phrase
        .split_whitespace()
        .flat_map(|word| word.split('-'))
        .filter(|word| word.contains(|c: char| c.is_alphabetic()))
        .map(|word| {
            word.chars()
                .find(|c| c.is_alphabetic())
                .unwrap()
                .to_ascii_uppercase()
        })
        .collect()
}

// fn is_camelcase(word: &str) -> bool {
//     word.chars().any(|c| c.is_uppercase())
// }
