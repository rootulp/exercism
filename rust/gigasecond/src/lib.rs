use time::Duration;
use time::PrimitiveDateTime as DateTime;

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    let gigasecond: Duration = Duration::SECOND * 1_000_000_000;
    return start + gigasecond;
}
