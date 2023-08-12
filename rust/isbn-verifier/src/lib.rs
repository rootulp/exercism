/// Determines whether the supplied string is a valid ISBN number
pub fn is_valid_isbn(input: &str) -> bool {
    if !is_valid_input(input) {
        return false;
    }
    let isbn = input
        .chars()
        .filter(|c| *c != '-')
        .enumerate()
        .map(|(i, c)| {
            if i == 9 && c == 'X' {
                10
            } else {
                c.to_digit(10).unwrap() as i32
            }
        })
        .collect::<Vec<i32>>();
    is_valid(isbn)
}

fn is_valid(isbn: Vec<i32>) -> bool {
    let sum = isbn[0] * 10
        + isbn[1] * 9
        + isbn[2] * 8
        + isbn[3] * 7
        + isbn[4] * 6
        + isbn[5] * 5
        + isbn[6] * 4
        + isbn[7] * 3
        + isbn[8] * 2
        + isbn[9];
    sum % 11 == 0
}

fn is_valid_input(input: &str) -> bool {
    let filtered: Vec<char> = input.chars().filter(|c| *c != '-').collect();
    if filtered.len() != 10 {
        return false;
    }
    if !is_valid_digits(filtered[0..9].to_vec()) {
        return false;
    }
    if !is_valid_check_character(filtered[9]) {
        return false;
    }
    true
}

fn is_valid_digits(digits: Vec<char>) -> bool {
    digits.iter().all(|c| c.is_ascii_digit())
}

fn is_valid_check_character(c: char) -> bool {
    c == 'X' || c.is_ascii_digit()
}
