use regex::Regex;

/// Check a Luhn checksum.
pub fn is_valid(input: &str) -> bool {
    let luhn = Luhn::new(input);

    match luhn {
        Ok(l) => l.is_valid(),
        Err(_) => false,
    }
}

struct Luhn {
    number: String,
}

impl Luhn {
    pub fn new(input: &str) -> Result<Self, &'static str> {
        let number = input.replace(" ", "");

        if number.len() == 1 {
            return Err("Invalid input");
        }

        let only_digits = Regex::new(r"^\d+$").unwrap();
        if !only_digits.is_match(number.as_str()) {
            return Err("Invalid input");
        }

        Ok(Luhn { number })
    }

    pub fn is_valid(&self) -> bool {
        unimplemented!()
    }
}
