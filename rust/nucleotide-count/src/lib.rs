use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if !is_valid_nucleotide(nucleotide) {
        return Err(nucleotide);
    }
    let counts = nucleotide_counts(dna)?;
    Ok(*counts.get(&nucleotide).unwrap())
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts = default_counts();
    for c in dna.chars() {
        if !is_valid_nucleotide(c) {
            return Err(c);
        }
        *counts.entry(c).or_insert(0) += 1;
    }
    Ok(counts)
}

fn default_counts() -> HashMap<char, usize> {
    HashMap::from([('A', 0), ('C', 0), ('G', 0), ('T', 0)])
}

fn is_valid_nucleotide(c: char) -> bool {
    matches!(c, 'A' | 'C' | 'G' | 'T')
}
