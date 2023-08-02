pub fn factors(n: u64) -> Vec<u64> {
    if n == 1 {
        return vec![];
    }
    let mut factors = vec![];
    let mut current = n;
    while !is_prime(current) {
        for prime in 2..=n {
            if current % prime == 0 {
                factors.push(prime);
                current /= prime;
                break;
            }
        }
    }
    factors.push(current);
    factors.sort();
    factors
}

fn is_prime(n: u64) -> bool {
    match n {
        0 | 1 => false,
        2 => true,
        _even if n % 2 == 0 => false,
        _ => {
            let sqrt_limit: u64 = (n as f64).sqrt() as u64;
            !(3..=sqrt_limit).step_by(2).any(|i| n % i == 0)
        }
    }
}
