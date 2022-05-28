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
        let check_sum = self.check_sum();
        println!("number: {}, check_sum: {}", self.number, check_sum);
        return self.check_sum() % 10 == 0;
    }

    fn check_sum(&self) -> u32 {
        let mut sum = 0;
        let digits = self.number.chars().rev();

        digits.enumerate().into_iter().for_each(|(i, digit)| {
            let mut digit = digit.to_digit(10).unwrap();
            if i % 2 == 1 {
                digit = if digit * 2 > 9 {
                    digit * 2 - 9
                } else {
                    digit * 2
                };
            }
            sum += digit;
        });

        sum
    }
}
