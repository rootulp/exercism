use std::collections::HashSet;

pub fn check(candidate: &str) -> bool {
    let mut seen: HashSet<char> = HashSet ::new();

    for char in candidate.to_lowercase().chars().filter(|c| c.is_alphabetic()) {
        let is_new = seen.insert(char);
        if !is_new {
            return false
        }
    }
    return true
}
