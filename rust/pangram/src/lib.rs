/// Determine whether a sentence is a pangram.
pub fn is_pangram(sentence: &str) -> bool {
    for letter in 'a'..='z' {
        if !sentence.to_lowercase().contains(letter) {
            return false;
        }
    }
    return true;
}
