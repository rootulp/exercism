pub fn collatz(n: u64) -> Option<u64> {
    if n == 0 {
        return None;
    }
    return Some(num_steps(n, 0));
}

fn num_steps(n: u64, steps: u64) -> u64 {
    if n == 1 {
        return steps
    }
    if is_even(n) {
        return num_steps(n / 2, steps + 1)
    }
    return num_steps((3 * n) + 1, steps + 1)
}

fn is_even(n: u64) -> bool {
    return n % 2 == 0
}
