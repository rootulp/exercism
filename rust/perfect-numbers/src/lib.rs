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
    println!("sum {}", sum);
    return if sum == num {
        Some(Classification::Perfect)
    } else if sum > num {
        Some(Classification::Abundant)
    } else {
        Some(Classification::Deficient)
    }
}

fn aliquot_sum(num: u64) -> u64 {
    factors(num).iter().sum()
}

fn factors(num: u64) -> Vec<u64> {
    let mut result = vec!();
    for i in 1..=num/2 {
        if num % i == 0 {
            result.push(i);
        }
    }
    println!("result {:?}", result);
    result
}
