pub fn abbreviate(phrase: &str) -> String {
    if phrase.is_empty() {
        return String::new();
    }
    phrase
        .split(' ')
        .map(|word| {
            word.chars()
                .find(|c| c.is_alphabetic())
                .unwrap_or(' ')
                .to_ascii_uppercase()
        })
        .filter(|c| c.is_alphabetic())
        .collect()
}
