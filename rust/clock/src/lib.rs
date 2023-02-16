use core::fmt;

#[derive(PartialEq, Eq, Debug)]
pub struct Clock {
    minutes: i32,
}

const HOURS_IN_A_DAY: i32 = 24;
const MINUTES_IN_AN_HOUR: i32 = 60;
const MINUTES_IN_A_DAY: i32 = HOURS_IN_A_DAY * MINUTES_IN_AN_HOUR;

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        Clock {
            minutes: normalize_minutes(hours * MINUTES_IN_AN_HOUR + minutes),
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self {
            minutes: normalize_minutes(self.minutes + minutes),
        }
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.minutes / 60, self.minutes % 60)
    }
}

fn normalize_minutes(minutes: i32) -> i32 {
    minutes.rem_euclid(MINUTES_IN_A_DAY)
}
