#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

pub struct BowlingGame {
    frames: Vec<Frame>,
    last_roll: Option<u16>
}

impl BowlingGame {
    pub fn new() -> Self {
        return BowlingGame{
            frames: vec!(),
            last_roll: None
        }
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if pins > 10 {
            return Err(Error::NotEnoughPinsLeft)
        }
        if pins == 10 {
            let frame = Frame::new(pins, None);
            self.frames.push(frame);
        }
        if self.last_roll.is_some() {
            let roll1 = self.last_roll.unwrap();
            let frame = Frame::new(roll1, Some(pins));
            self.frames.push(frame);
            self.last_roll = None;
        } else {
            self.last_roll = Some(pins);
        }
        Ok(())
    }

    pub fn score(&self) -> Option<u16> {
        if !self.is_game_over() {
            return None
        }
        let mut total = 0;
        for frame in &self.frames {
            total += frame.score();
        }
        return Some(total)
    }

    fn is_game_over(&self) -> bool {
        self.frames.len() == 10
    }
}

pub struct Frame {
    roll1: u16,
    roll2: Option<u16>
}

impl Frame {
    pub fn new(roll1:u16, roll2: Option<u16>) -> Self {
        return Frame{
            roll1: roll1,
            roll2: roll2
        }
    }

    pub fn score(&self) -> u16 {
        if self.is_strike() {
            todo!("not implemented")
        }
        if self.is_spare() {
            todo!("not implemented")
        }
        self.open_frame_score()
    }

    fn is_strike(&self) -> bool {
        return self.roll1 == 10
    }
    fn is_spare(&self) -> bool {
        return self.roll1 != 10 && self.open_frame_score() == 10
    }
    fn open_frame_score(&self) -> u16 {
        self.roll1 + self.roll2.unwrap_or(0)
    }
}
