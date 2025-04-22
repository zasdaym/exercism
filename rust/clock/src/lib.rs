use std::fmt;

const MINUTES_IN_A_DAY: i32 = 24 * 60;

#[derive(Debug, Eq, PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let total_minutes = hours * 60 + minutes;
        let normalized_minutes =
            ((total_minutes % MINUTES_IN_A_DAY) + MINUTES_IN_A_DAY) % (MINUTES_IN_A_DAY);

        Clock {
            minutes: normalized_minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(0, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.minutes / 60, self.minutes % 60)
    }
}
