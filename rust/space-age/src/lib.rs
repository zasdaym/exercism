#[derive(Debug)]
pub struct Duration {
    earth_years: f64,
}

const EARTH_YEAR_IN_SECONDS: f64 = 31557600.0;

impl From<u64> for Duration {
    fn from(seconds: u64) -> Self {
        Self {
            earth_years: (seconds as f64) / EARTH_YEAR_IN_SECONDS,
        }
    }
}

pub trait Planet {
    const EARTH_YEARS_RATIO: f64;
    fn years_during(d: &Duration) -> f64 {
        d.earth_years / Self::EARTH_YEARS_RATIO
    }
}

macro_rules! planet {
    ($n:ident, $p:expr) => {
        pub struct $n;
        impl Planet for $n {
            const EARTH_YEARS_RATIO: f64 = $p;
        }
    };
}

planet!(Mercury, 0.2408467);
planet!(Venus, 0.61519726);
planet!(Earth, 1.0);
planet!(Mars, 1.8808158);
planet!(Jupiter, 11.862615);
planet!(Saturn, 29.447498);
planet!(Uranus, 84.016846);
planet!(Neptune, 164.79132);
