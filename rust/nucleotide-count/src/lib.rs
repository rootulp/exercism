use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if !valid_nucleotide(nucleotide) {
        return Err(nucleotide)
    }
    let counts = nucleotide_counts(dna)?;
    let result = counts.get(&nucleotide).unwrap_or(&0);
    return Ok(*result)
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts: HashMap<char, usize> = default_map();
    for c in dna.chars() {
        if !valid_nucleotide(c) {
            return Err(c)
        }
        let default: usize = 0;
        let v = counts.get(&c).unwrap_or(&default);
        counts.insert(c, v + 1);
    }
    return Ok(counts)
}

fn default_map() -> HashMap<char, usize> {
    let result = HashMap::from([
        ('A', 0),
        ('C', 0),
        ('G', 0),
        ('T', 0),
    ]);
    result
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
