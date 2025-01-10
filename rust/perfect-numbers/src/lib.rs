#[derive(Debug, PartialEq, Eq)]
pub enum Classification {
    Abundant,
    Perfect,
    Deficient,
}

pub fn classify(num: u64) -> Option<Classification> {
    if num == 0 {
        return None;
    }

    let sum = aliquot_sum(num);
    return if sum == num {
        Some(Classification::Perfect)
    } else if sum > num {
        Some(Classification::Abundant)
    } else {
        Some(Classification::Deficient)
    }
}

fn aliquot_sum(num: u64) -> u64 {
    get_factors(num).iter().sum()
}

fn get_factors(num: u64) -> Vec<u64> {
    let mut factors = vec!();
    for x in 1..=num/2 {
        if num % x == 0 {
            factors.push(x);
        }
    }
    factors
}
