use rand::Rng;

pub fn private_key(prime: u64) -> u64 {
    let mut rng = rand::thread_rng();
    let result = rng.gen_range(2..prime);
    return result;
}

pub fn public_key(p: u64, g: u64, a: u64) -> u64 {
    todo!("Calculate public key using prime numbers {p} and {g}, and private key {a}")
}

pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    todo!("Calculate secret key using prime number {p}, public key {b_pub}, and private key {a}")
}
