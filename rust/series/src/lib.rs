pub fn series(digits: &str, len: usize) -> Vec<String> {
    let windows: Vec<&str> = (0..=digits.len() - len)
    .map(|i| &digits[i..i + len])
    .collect();
    windows.iter().map(|&s| s.to_string()).collect()
}
