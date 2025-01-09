use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    let mut seen: HashSet<char> = HashSet ::new();
    for char in candidate.to_lowercase().chars() {
        if !char.is_alphabetic() {
            continue
        }
        if seen.contains(&char) {
            return false
        }
        seen.insert(char);
    }
    return true
}
