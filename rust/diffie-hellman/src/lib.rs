use rand::Rng;

pub fn private_key(p: u64) -> u64 {
    let mut rng = rand::thread_rng();
    let result = rng.gen_range(2..p);
    return result;
}

pub fn public_key(p: u64, g: u64, a: u64) -> u64 {
    // A = gᵃ mod p
    return g.pow(a.try_into().unwrap()) % p
}

pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    todo!("Calculate secret key using prime number {p}, public key {b_pub}, and private key {a}")
}
