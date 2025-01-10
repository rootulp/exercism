use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if !valid_nucleotide(nucleotide) {
        return Err(nucleotide);
    }
    let counts = nucleotide_counts(dna)?;
    Ok(*counts.get(&nucleotide).unwrap())
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    for c in dna.chars() {
        if !valid_nucleotide(c) {
            return Err(c);
        }
    }
    let mut counts: HashMap<char, usize> = default_map();
    for c in dna.chars() {
        let default: usize = 0;
        let v = counts.get(&c).unwrap_or(&default);
        counts.insert(c, v + 1);
    }
    Ok(counts)
}

fn default_map() -> HashMap<char, usize> {
    HashMap::from([('A', 0), ('C', 0), ('G', 0), ('T', 0)])
}

fn valid_nucleotide(c: char) -> bool {
    matches!(c, 'A' | 'C' | 'G' | 'T')
}
