use std::collections::HashSet;

pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    unique_multiples(limit, factors).iter().sum()
}

fn unique_multiples(limit: u32, factors: &[u32]) -> Vec<u32> {
    let mut result = HashSet::new();
    for factor in factors {
        result.extend(multiples(limit, *factor));
    }
    result.into_iter().collect()
}

fn multiples(limit: u32, factor: u32) -> Vec<u32> {
    if factor == 0 {
        return vec![];
    }
    (1..limit).filter(|x| x % factor == 0).collect()
}
