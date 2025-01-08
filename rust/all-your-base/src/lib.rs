#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(u32),
}

///
/// Convert a number between two bases.
///
/// A number is any slice of digits.
/// A digit is any unsigned integer (e.g. u8, u16, u32, u64, or usize).
/// Bases are specified as unsigned integers.
///
/// Return the corresponding Error enum if the conversion is impossible.
///
///
/// You are allowed to change the function signature as long as all test still pass.
///
///
/// Example:
/// Input
///   number: &[4, 2]
///   from_base: 10
///   to_base: 2
/// Result
///   Ok(vec![1, 0, 1, 0, 1, 0])
///
/// The example corresponds to converting the number 42 from decimal
/// which is equivalent to 101010 in binary.
///
///
/// Notes:
///  * The empty slice ( "[]" ) is equal to the number 0.
///  * Never output leading 0 digits, unless the input number is 0, in which the output must be `[0]`.
///    However, your function must be able to process input with leading 0 digits.
///
pub fn convert(number: &[u32], from_base: u32, to_base: u32) -> Result<Vec<u32>, Error> {
    if !is_valid_base(from_base) {
        return Err(Error::InvalidInputBase)
    }
    if !is_valid_base(to_base) {
        return Err(Error::InvalidOutputBase)
    }
    validate_digit(number, from_base)?;

    if number.is_empty() {
        return Ok(vec!(0))
    }
    if number.iter().all(|digit| *digit == 0) {
        return Ok(vec!(0))
    }

    let decimal = convert_from(number, from_base);
    let converted = convert_to(decimal, to_base);
    return Ok(converted)
}

// convert_to converts decimal to to_base
fn convert_to(decimal: u32, to_base: u32) -> Vec<u32> {
    let mut result = vec!();
    let mut n = decimal;
    while n != 0 {
        let remainder = n % to_base;
        n = n / to_base;
        result.insert(0, remainder);
    }
    return result
}

// convert_from converts number from_base to decimal
fn convert_from(number: &[u32], from_base: u32) -> u32 {
    let mut decimal = 0;
    let mut cloned = number.to_vec();
    cloned.reverse();
    for (i, digit) in cloned.iter().enumerate() {
        decimal += digit * from_base.pow(i.try_into().unwrap());
    }
    decimal
}

fn is_valid_base(base :u32) -> bool {
    if base == 0 || base == 1 {
        return false
    }
    return true
}

fn validate_digit(number: &[u32], from_base: u32) -> Result<u32, Error> {
    for digit in number {
        if *digit >= from_base {
            return Err(Error::InvalidDigit(*digit))
        }
    }
    Ok(0)
}
