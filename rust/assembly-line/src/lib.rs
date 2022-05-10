pub fn production_rate_per_hour(speed: u8) -> f64 {
    const CARS_PER_HOUR: u32 = 221;
    let production = CARS_PER_HOUR * speed as u32;
    let success_rate = success_rate(speed);
    success_rate * production as f64
}

pub fn success_rate(speed: u8) -> f64 {
    if (1..=4).contains(&speed) {
        return 1.0;
    }
    if (5..=8).contains(&speed) {
        return 0.90;
    }
    if (9..=10).contains(&speed) {
        return 0.77;
    }
    // Ideally this would throw an error
    0.0
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
}
