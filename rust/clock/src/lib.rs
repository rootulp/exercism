use std::fmt;

#[derive(PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        return Clock {
            minutes: (hours * 60) + minutes,
        };
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        return Clock::new(0, self.minutes + minutes);
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut minutes = self.minutes;
        let mut hours = 0;

        while minutes < 0 {
            hours -= 1;
            minutes += 60
        }
        while minutes >= 60 {
            hours += 1;
            minutes -= 60
        }
        while hours < 0 {
            hours += 24;
        }
        while hours >= 24 {
            hours -= 24;
        }

        write!(f, "{:02}:{:02}", hours, minutes)
    }
}

impl fmt::Debug for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.minutes)
    }
}
