#[derive(Debug)]
pub struct Duration {
    seconds: u64,
}

impl From<u64> for Duration {
    fn from(seconds: u64) -> Self {
        return Self { seconds };
    }
}

pub trait Planet {
    fn years_during(duration: &Duration) -> f64 {
        unimplemented!(
            "convert a duration ({:?}) to the number of years on this planet for that duration",
            duration,
        );
    }
}

pub struct Mercury;
pub struct Venus;
pub struct Earth;
pub struct Mars;
pub struct Jupiter;
pub struct Saturn;
pub struct Uranus;
pub struct Neptune;

impl Planet for Mercury {
    fn years_during(duration: &Duration) -> f64 {
        const MERCURY_ORBITAL_PERIOD: f64 = 0.2408467;
        return Earth::years_during(duration) / MERCURY_ORBITAL_PERIOD;
    }
}
impl Planet for Venus {}
impl Planet for Earth {
    fn years_during(duration: &Duration) -> f64 {
        const SECONDS_IN_EARTH_YEAR: f64 = 31557600.0;
        return duration.seconds as f64 / SECONDS_IN_EARTH_YEAR;
    }
}
impl Planet for Mars {}
impl Planet for Jupiter {}
impl Planet for Saturn {}
impl Planet for Uranus {}
impl Planet for Neptune {}
