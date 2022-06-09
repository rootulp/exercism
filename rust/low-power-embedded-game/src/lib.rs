pub fn divmod(dividend: i16, divisor: i16) -> (i16, i16) {
    (dividend / divisor, dividend % divisor)
}

pub fn evens<T>(iter: impl Iterator<Item = T>) -> impl Iterator<Item = T> {
    iter.enumerate()
        .filter(|(index, _)| index % 2 == 0)
        .map(|(_, val)| val)
}

pub struct Position(pub i16, pub i16);
impl Position {
    pub fn manhattan(&self) -> i16 {
        let delta_x = (&self.0).abs();
        let delta_y = (&self.1).abs();
        delta_x + delta_y
    }
}
