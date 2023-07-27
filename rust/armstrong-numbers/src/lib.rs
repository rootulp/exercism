pub fn is_armstrong_number(num: u32) -> bool {
    let string: String = num.to_string();
    let digits = string.chars().map(|d| d.to_digit(10).unwrap());
    let num_digits = string.len().try_into().unwrap();
    let powered_sum: u64 = digits
        .map(|digit: u32| (digit as u64).pow(num_digits))
        .sum();
    num as u64 == powered_sum
}
