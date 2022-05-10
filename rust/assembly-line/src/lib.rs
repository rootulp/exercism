pub fn production_rate_per_hour(speed: u8) -> f64 {
    const CARS_PER_HOUR: u32 = 221;
    let production = CARS_PER_HOUR * speed as u32;
    let success_rate = success_rate(speed);
    let result = success_rate * production as f64;
    println!("result {result} for speed {speed}");
    return result
    // 1989 * .70 = 1392.3
}

pub fn success_rate(speed: u8) -> f64 {
    if speed >= 1 && speed <= 4 {
        return 1.0;
    }
    if speed >= 5 && speed <= 8 {
        return 0.90;
    }
    if speed >= 9 && speed <= 10 {
        return 0.77;
    }
    // Ideally this would throw an error
    return 0.0;
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    unimplemented!("calculate the amount of working items at speed: {}", speed)
}
