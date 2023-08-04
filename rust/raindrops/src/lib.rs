pub fn raindrops(n: u32) -> String {
    if is_factor(n, 3) {
        return "Pling".to_string();
    } else if is_factor(n, 5) {
        return "Plang".to_string();
    } else if is_factor(n, 7) {
        return "Plong".to_string();
    }
    n.to_string()
}

fn is_factor(n: u32, factor: u32) -> bool {
    n % factor == 0
}
