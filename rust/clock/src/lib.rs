use std::fmt;

#[derive(PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        return Clock { hours, minutes };
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        return Clock::new(self.hours, self.minutes + minutes);
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}:{})", self.hours, self.minutes)
    }
}

impl fmt::Debug for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}:{})", self.hours, self.minutes)
    }
}
