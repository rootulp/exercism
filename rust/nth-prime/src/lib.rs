pub fn nth(n: u32) -> u32 {
    let n_usize: usize = n as usize;
    let mut primes = Vec::new();
    let mut candidate = 2;
    while primes.len() != n_usize + 1 {
        if is_prime(candidate) {
            primes.push(candidate);
        }
        candidate += 1;
    }
    primes[n_usize]
}

fn is_prime(n: u32) -> bool {
    match n {
        0 | 1 => false,
        2 => true,
        _even if n % 2 == 0 => false,
        _ => {
            let limit = (n as f64).sqrt() as u32;
            !(3..=limit).step_by(2).any(|i| n % i == 0)
        }
    }
}
