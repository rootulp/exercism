pub fn egg_count(decimal: u32) -> usize {
    let binary = decimal_to_binary(decimal);
    count('1', binary)
}

fn decimal_to_binary(decimal: u32) -> String {
    let mut result = "".to_string();
    let mut n = decimal;
    while n != 0 {
        let remainder = n % 2;
        n /= 2;
        result.insert(0, char::from_digit(remainder, 10).unwrap());
    }
    result
}

fn count(c: char, string: String) -> usize {
    string.chars().filter(|char| *char == c).count()
}
