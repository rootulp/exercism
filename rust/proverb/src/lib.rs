pub fn build_proverb(list: &[&str]) -> String {
    if list.is_empty() {
        return String::new();
    }
    let mut result = String::new();
    for pair in list.windows(2) {
        result += &format!("For want of a {} the {} was lost.\n", pair[0], pair[1])
    }
    result += &format!("And all for the want of a {}.", list[0]);
    result
}
