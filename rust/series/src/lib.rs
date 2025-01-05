pub fn series(digits: &str, len: usize) -> Vec<String> {
    if len > digits.len() {
        return vec![];
    }

    (0..=digits.len() - len)
        .map(|i| &digits[i..i + len])
        .collect::<Vec<&str>>()
        .iter()
        .map(|&s| s.to_string())
        .collect()
}
