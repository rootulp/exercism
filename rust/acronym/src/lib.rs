pub fn abbreviate(phrase: &str) -> String {
    if phrase.is_empty() {
        return String::new();
    }
    phrase
        .split(' ')
        .map(|word| {
            word.chars()
                .find(|c| c.is_alphabetic())
                .unwrap()
                .to_ascii_uppercase()
        })
        .collect()
}
