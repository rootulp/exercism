use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    let mut seen: HashSet<char> = HashSet ::new();

    for char in candidate.to_lowercase().chars().filter(|c| c.is_alphabetic()) {
        if seen.contains(&char) {
            return false
        }
        seen.insert(char);
    }
    return true
}
