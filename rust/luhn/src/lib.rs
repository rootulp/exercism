/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let luhn = Luhn::new(code);

    match luhn {
        Ok(l) => l.is_valid(),
        Err(_) => false,
    }
}

struct Luhn {
    code: String,
}

impl Luhn {
    pub fn new(code: &str) -> Result<Self, &'static str> {
        if code.len() == 1 {
            return Err("Invalid code");
        }
        Ok(Luhn {
            code: code.to_string(),
        })
    }

    pub fn is_valid(&self) -> bool {
        unimplemented!()
    }
}
