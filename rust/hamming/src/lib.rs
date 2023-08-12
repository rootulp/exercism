/// Return the Hamming distance between the strings,
/// or None if the lengths are mismatched.
pub fn hamming_distance(s1: &str, s2: &str) -> Option<usize> {
    if s1.len() != s2.len() {
        return None;
    }
    let mut distance = 0;
    for i in 0..s1.len() {
        if s1.chars().nth(i) != s2.chars().nth(i) {
            distance += 1;
        }
    }
    Some(distance)
}
