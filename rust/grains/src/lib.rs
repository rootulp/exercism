pub fn square(s: u32) -> u64 {
    if !(1..=64).contains(&s) {
        panic!("Square must be between 1 and 64");
    }
    if s == 1 {
        return 1;
    }
    square(s - 1) * 2
}

pub fn total() -> u64 {
    unimplemented!();
}
