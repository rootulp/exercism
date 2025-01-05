pub fn series(digits: &str, len: usize) -> Vec<String> {
    if len > digits.len() {
        return vec![];
    }

    let windows: Vec<&str> = (0..=digits.len() - len)
    .map(|i| &digits[i..i + len])
    .collect();
    windows.iter().map(|&s| s.to_string()).collect()
}
