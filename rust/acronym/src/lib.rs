pub fn abbreviate(phrase: &str) -> String {
    phrase
        .split_whitespace()
        .flat_map(|word| word.split('-')) // split on hyphens
        .flat_map(split_camel_or_pascal_case) // split on camelCase or PascalCase
        .filter(|word| word.contains(char::is_alphabetic))
        .map(|word| {
            word.chars()
                .find(|c| c.is_alphabetic())
                .unwrap()
                .to_ascii_uppercase()
        })
        .collect()
}

fn split_camel_or_pascal_case(s: &str) -> Vec<String> {
    let mut result = Vec::new();
    let mut word_start = 0;

    for (i, (prev_char, curr_char)) in s.chars().zip(s.chars().skip(1)).enumerate() {
        if prev_char.is_lowercase() && curr_char.is_uppercase() {
            result.push(s[word_start..=i].to_string());
            word_start = i + 1;
        }
    }

    // Push the last word
    result.push(s[word_start..].to_string());
    result
}
