const CARS_PER_HOUR: f64 = 221.0;

fn success_rate(speed: u8) -> f64 {
    return match speed {
        1..=4 => 1.0,
        5..=8 => 0.9,
        9..=10 => 0.77,
        _ => 0.0,
    };
}

pub fn production_rate_per_hour(speed: u8) -> f64 {
    return speed as f64 * CARS_PER_HOUR * success_rate(speed);
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    return production_rate_per_hour(speed) as u32 / 60;
}
