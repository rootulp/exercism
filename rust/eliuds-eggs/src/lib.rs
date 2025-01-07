pub fn egg_count(decimal: u32) -> usize {
    let binary = decimal_to_binary(decimal);
    let eggs = count('1', binary);
    eggs
}

fn decimal_to_binary(decimal: u32) -> String {
    let mut result = "".to_string();

    let mut n = decimal;
    while n != 0 {
        let remainder = n % 2;
        n = n / 2;
        let c = char::from_digit(remainder, 10).unwrap();
        result.insert(0, c);
    }

    result
}

fn count(c: char, string: String) -> usize {
    string.chars().filter(|char| *char == c).count()
}
