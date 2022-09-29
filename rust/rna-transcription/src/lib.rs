use std::collections::HashMap;

#[derive(Debug, PartialEq, Eq)]
pub struct Dna<'a> {
    contents: &'a str,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Rna<'a> {
    contents: &'a str,
}

impl<'a> Dna<'a> {
    const VALID_NUCLEOTIDES: &str = "ACGT";

    pub fn new(dna: &str) -> Result<Dna, usize> {
        for (i, c) in dna.chars().enumerate() {
            if !Self::VALID_NUCLEOTIDES.contains(c) {
                return Err(i);
            }
        }
        Ok(Dna { contents: dna })
    }

    pub fn into_rna(self) -> Rna<'a> {
        let complements = HashMap::from([('G', 'C'), ('C', 'G'), ('T', 'A'), ('A', 'U')]);

        let mut rna_contents = String::new();
        for (_, c) in self.contents.chars().enumerate() {
            rna_contents.push(*complements.get(&c).unwrap());
        }

        let rna = Rna::new(rna_contents.as_str()).expect("dna could not be converted to rna");
        // can't return the following rna b/c Rust borrow checker
        rna
    }
}

impl<'a> Rna<'a> {
    const VALID_NUCLEOTIDES: &str = "ACGU";
    pub fn new(rna: &str) -> Result<Rna, usize> {
        for (i, c) in rna.chars().enumerate() {
            if !Self::VALID_NUCLEOTIDES.contains(c) {
                return Err(i);
            }
        }
        Ok(Rna { contents: rna })
    }
}
