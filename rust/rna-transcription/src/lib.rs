use std::collections::HashMap;

#[derive(Debug, PartialEq, Eq)]
pub struct Dna {
    contents: String,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Rna {
    contents: String,
}

impl Dna {
    const VALID_NUCLEOTIDES: &str = "ACGT";

    pub fn new(input: &str) -> Result<Dna, usize> {
        for (i, c) in input.chars().enumerate() {
            if !Self::VALID_NUCLEOTIDES.contains(c) {
                return Err(i);
            }
        }
        Ok(Dna {
            contents: input.to_string(),
        })
    }

    pub fn into_rna(self) -> Rna {
        let complements = HashMap::from([('G', 'C'), ('C', 'G'), ('T', 'A'), ('A', 'U')]);

        let mut converted = String::new();

        for (_, c) in self.contents.chars().enumerate() {
            converted.push(*complements.get(&c).unwrap())
        }

        let result = Rna::new(converted.as_str()).expect("dna could not be converted to rna");
        result
    }
}

impl Rna {
    const VALID_NUCLEOTIDES: &str = "ACGU";
    pub fn new(input: &str) -> Result<Rna, usize> {
        for (i, c) in input.chars().enumerate() {
            if !Self::VALID_NUCLEOTIDES.contains(c) {
                return Err(i);
            }
        }
        Ok(Rna {
            contents: input.to_string(),
        })
    }
}
