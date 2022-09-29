#[derive(Debug, PartialEq, Eq)]
pub struct Dna<'a> {
    contents: &'a str,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Rna;

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

    pub fn into_rna(self) -> Rna {
        unimplemented!("Transform Dna {:?} into corresponding Rna", self);
    }
}

impl Rna {
    pub fn new(rna: &str) -> Result<Rna, usize> {
        unimplemented!("Construct new Rna from '{}' string. If string contains invalid nucleotides return index of first invalid nucleotide", rna);
    }
}
