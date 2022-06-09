#[derive(Debug)]
pub struct Duration {
    seconds: u64,
}

impl From<u64> for Duration {
    fn from(seconds: u64) -> Self {
        Self { seconds }
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
        Earth::years_during(duration) / MERCURY_ORBITAL_PERIOD
    }
}
impl Planet for Venus {
    fn years_during(duration: &Duration) -> f64 {
        const VENUS_ORBITAL_PERIOD: f64 = 0.61519726;
        Earth::years_during(duration) / VENUS_ORBITAL_PERIOD
    }
}
impl Planet for Earth {
    fn years_during(duration: &Duration) -> f64 {
        const SECONDS_IN_EARTH_YEAR: f64 = 31557600.0;
        duration.seconds as f64 / SECONDS_IN_EARTH_YEAR
    }
}
impl Planet for Mars {
    fn years_during(duration: &Duration) -> f64 {
        const MARS_ORBITAL_PERIOD: f64 = 1.8808158;
        Earth::years_during(duration) / MARS_ORBITAL_PERIOD
    }
}
impl Planet for Jupiter {
    fn years_during(duration: &Duration) -> f64 {
        const JUPITER_ORBITAL_PERIOD: f64 = 11.862615;
        Earth::years_during(duration) / JUPITER_ORBITAL_PERIOD
    }
}
impl Planet for Saturn {
    fn years_during(duration: &Duration) -> f64 {
        const SATURN_ORBITAL_PERIOD: f64 = 29.447498;
        Earth::years_during(duration) / SATURN_ORBITAL_PERIOD
    }
}
impl Planet for Uranus {
    fn years_during(duration: &Duration) -> f64 {
        const URANUS_ORBITAL_PERIOD: f64 = 84.016846;
        Earth::years_during(duration) / URANUS_ORBITAL_PERIOD
    }
}
impl Planet for Neptune {
    fn years_during(duration: &Duration) -> f64 {
        const NEPTUNE_ORBITAL_PERIOD: f64 = 164.79132;
        Earth::years_during(duration) / NEPTUNE_ORBITAL_PERIOD
    }
}
