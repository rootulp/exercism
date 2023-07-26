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

        let mut rna = self.contents.chars().collect::<Vec<_>>();
        for (i, c) in self.contents.chars().enumerate() {
            rna[i] = *complements.get(&c).unwrap();
        }

        let result: Rna<'_> = Rna::new(self.contents).expect("dna could not be converted to rna");
        result
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
