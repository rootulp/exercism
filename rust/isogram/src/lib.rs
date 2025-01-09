use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    if candidate.is_empty() {
        return true
    }
    let mut seen: HashSet<char> = HashSet ::new();
    for char in candidate.to_lowercase().chars() {
        if seen.contains(&char) {
            return false
        }
        seen.insert(char);
    }
    return true
}
