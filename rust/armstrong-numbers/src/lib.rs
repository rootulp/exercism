pub fn is_armstrong_number(num: u32) -> bool {
    let string = num.to_string();
    let digits = string.chars().map(|d| d.to_digit(10).unwrap());
    let num_digits = string.len().try_into().unwrap();
    num == digits.map(|digit: u32| digit.pow(num_digits)).sum()
}
