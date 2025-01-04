use num_bigint::ToBigUint;
use rand::Rng;

pub fn private_key(p: u64) -> u64 {
    let mut rng = rand::thread_rng();
    rng.gen_range(2..p)
}

pub fn public_key(p: u64, g: u64, a: u64) -> u64 {
    g.to_biguint().unwrap().modpow(&a.to_biguint().unwrap(), &p.to_biguint().unwrap()).try_into().unwrap()
}

pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    b_pub.to_biguint().unwrap().modpow(&a.to_biguint().unwrap(), &p.to_biguint().unwrap()).try_into().unwrap()
}
