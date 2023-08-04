pub fn raindrops(n: u32) -> String {
    let mut words = String::new();
    if n % 3 == 0 {
        words.push_str("Pling")
    }
    if n % 5 == 0 {
        words.push_str("Plang")
    }
    if n % 7 == 0 {
        words.push_str("Plong")
    }
    if words.is_empty() {
        return n.to_string();
    }
    words
}
