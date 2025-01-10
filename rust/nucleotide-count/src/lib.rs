use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if !valid_nucleotide(nucleotide) {
        return Err(nucleotide)
    }
    let mut occurences: HashMap<char, usize> = HashMap::new();
    for c in dna.chars() {
        if !valid_nucleotide(c) {
            return Err(c)
        }
        let default: usize = 0;
        let v = occurences.get(&c).unwrap_or(&default);
        occurences.insert(c, v + 1);
    }
    let result = occurences.get(&nucleotide).unwrap_or(&0);
    return Ok(*result)
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
