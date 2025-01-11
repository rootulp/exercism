#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

pub struct BowlingGame {
}

impl BowlingGame {
    pub fn new() -> Self {
        return BowlingGame{}
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        Ok(())
    }

    pub fn score(&self) -> Option<u16> {
        todo!("Return the score if the game is complete, or None if not.");
    }
}
