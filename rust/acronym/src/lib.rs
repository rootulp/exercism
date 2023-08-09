pub fn abbreviate(phrase: &str) -> String {
    if phrase.is_empty() {
        return String::new();
    }
    phrase
        .split_whitespace()
        .flat_map(|word| word.split('-'))
        // .flat_map(|word| word.split(|c: char| is_pascal_case(word) && c.is_()))
        .flat_map(split_camel_or_pascal_case)
        .filter(|word| word.contains(|c: char| c.is_alphabetic()))
        .map(|word| {
            word.chars()
                .find(|c| c.is_alphabetic())
                .unwrap()
                .to_ascii_uppercase()
        })
        .collect()
}

// fn is_pascal_case(word: &str) -> bool {
//     if word.is_empty() {
//         return false;
//     }
//     if word
//         .chars()
//         .filter(|c| c.is_alphabetic())
//         .all(|c| c.is_uppercase())
//     {
//         return false;
//     }
//     if word.chars().filter(|c| c.is_uppercase()).count() <= 1 {
//         return false;
//     }
//     true
// }

fn split_camel_or_pascal_case(s: &str) -> Vec<String> {
    let mut result = Vec::new();
    let mut word_start = 0;

    for (index, (prev_char, curr_char)) in s.chars().zip(s.chars().skip(1)).enumerate() {
        if prev_char.is_lowercase() && curr_char.is_uppercase() {
            result.push(s[word_start..=index].to_string());
            word_start = index + 1;
        }
    }

    // Push the last word
    result.push(s[word_start..].to_string());

    result
}
