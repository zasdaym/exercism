#[derive(Debug)]
pub struct Duration {
    seconds: f64,
}

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Duration { seconds: s as f64 }
    }
}

const EARTH_YEAR_IN_SECONDS: f64 = 365.25 * 24.0 * 60.0 * 60.0;

pub trait Planet {
    const ORBITAL_PERIOD: f64;

    fn years_during(d: &Duration) -> f64 {
        d.seconds / EARTH_YEAR_IN_SECONDS / Self::ORBITAL_PERIOD
    }
}

macro_rules! planet {
    ($($name:ident, $period:expr);*) => {
        $(
            pub struct $name;
            impl Planet for $name {
                const ORBITAL_PERIOD: f64 = $period;
            }
        )*
    };
}

planet! {
    Mercury, 0.2408467;
    Venus, 0.61519726;
    Earth, 1.0;
    Mars, 1.8808158;
    Jupiter, 11.862615;
    Saturn, 29.447498;
    Uranus, 84.016846;
    Neptune, 164.79132
}
