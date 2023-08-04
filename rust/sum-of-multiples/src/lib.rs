use std::collections::HashSet;

pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let mut points = HashSet::new();

    for factor in factors {
        let multiples = multiples(limit, *factor);
        for multiple in multiples {
            points.insert(multiple);
        }
    }
    points.into_iter().sum()
}

fn multiples(limit: u32, factor: u32) -> Vec<u32> {
    if factor == 0 {
        return vec![];
    }
    (1..limit).filter(|x| x % factor == 0).collect()
}
