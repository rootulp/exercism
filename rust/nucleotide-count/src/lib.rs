use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if !valid_nucleotide(nucleotide) {
        return Err(nucleotide)
    }
    for c in dna.chars() {
        if !valid_nucleotide(c) {
            return Err(c)
        }
    }
    Ok(0)
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    todo!("How much of every nucleotide type is contained inside DNA string '{dna}'?");
}

fn valid_nucleotide(nucleotide: char) -> bool {
    return match nucleotide {
        'A' => true,
        'C' => true,
        'G' => true,
        'T' => true,
        _ => false,
    }
}
