use num_bigint::ToBigUint;
use rand::Rng;

pub fn private_key(p: u64) -> u64 {
    let mut rng = rand::thread_rng();
    let result = rng.gen_range(2..p);
    return result;
}

pub fn public_key(p: u64, g: u64, a: u64) -> u64 {
    let pbig = p.to_biguint().unwrap();
    let gbig = g.to_biguint().unwrap();
    let abig = a.to_biguint().unwrap();
    let result = gbig.modpow(&abig, &pbig);
    return result.try_into().unwrap();
}

pub fn secret(p: u64, b_pub: u64, a: u64) -> u64 {
    return b_pub.pow(a.try_into().unwrap()) % p
}
