use std::collections::HashSet;

pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    unique_multiples(limit, factors).iter().sum()
}

fn unique_multiples(limit: u32, factors: &[u32]) -> Vec<u32> {
    let multiples: Vec<_> = factors.iter().flat_map(|x| multiples(limit, *x)).collect();
    let unique_multiples: HashSet<u32> = multiples.into_iter().collect();
    unique_multiples.into_iter().collect()
}

fn multiples(limit: u32, factor: u32) -> Vec<u32> {
    if factor == 0 {
        return vec![];
    }
    (1..limit).filter(|x| x % factor == 0).collect()
}
