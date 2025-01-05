pub fn series(digits: &str, len: usize) -> Vec<String> {
    if len == 1 {
        return digits.chars().map(|c| c.to_string()).collect();
    }
    return vec![];
}
