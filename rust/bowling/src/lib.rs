#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Frame {
    pub roll1: u16,
    pub roll2: u16
}

impl Frame {
    pub fn new(roll1: u16, roll2: u16) -> Self {
        Self { roll1, roll2 }
    }

    pub fn score(&self, next_roll1: Option<u16>, next_roll2: Option<u16>) -> u16 {
        if self.is_strike() {
            10 + next_roll1.unwrap_or(0) + next_roll2.unwrap_or(0)
        } else if self.is_spare() {
            10 + next_roll1.unwrap_or(0)
        } else {
            self.open_frame_score()
        }
    }

    fn is_strike(&self) -> bool {
        return self.roll1 == 10
    }
    fn is_spare(&self) -> bool {
        return self.roll1 != 10 && (self.roll1 + self.roll2) == 10
    }
    fn open_frame_score(&self) -> u16 {
        self.roll1 + self.roll2
    }
}

pub struct BowlingGame {
    frames: Vec<Frame>,
    /// This holds a single roll that hasn't yet formed a complete Frame (roll1 but not roll2).
    pending_roll: Option<u16>,
    /// Extra fill balls in the 10th frame scenario
    fill_ball1: Option<u16>,
    fill_ball2: Option<u16>
}

impl BowlingGame {
    pub fn new() -> Self {
        return BowlingGame{
            frames: vec!(),
            pending_roll: None,
            fill_ball1: None,
            fill_ball2: None,
        }
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if self.is_game_over() {
            return Err(Error::GameComplete)
        }
        if pins > 10 {
            return Err(Error::NotEnoughPinsLeft);
        }
        if let Some(prev) = self.pending_roll {
            if prev + pins > 10 {
                return Err(Error::NotEnoughPinsLeft);
            }
        }

        // Handle fill balls
        if self.frames.len() == 10 {
            if let Some(last) = self.frames.last() {
                if last.is_spare() {
                    self.fill_ball1 = Some(pins);
                } else if last.is_strike() {
                    if self.fill_ball1.is_none() {
                        self.fill_ball1 = Some(pins);
                    } else {
                        if self.fill_ball1.unwrap() != 10 && self.fill_ball1.unwrap() + pins > 10 {
                            return Err(Error::NotEnoughPinsLeft)
                        }
                        self.fill_ball2 = Some(pins);
                    }
                }
            }
            return Ok(())
        }

        // Handle normal rolls
        if pins == 10 {
            // Handle strike
            let frame = Frame::new(pins, 0);
            self.frames.push(frame);
            return Ok(())
        }
        if let Some(prev) = self.pending_roll.take() {
            let frame = Frame::new(prev, pins);
            self.frames.push(frame);
            return Ok(())
        } else {
            self.pending_roll = Some(pins);
            return Ok(())
        }
    }

    pub fn score(&self) -> Option<u16> {
        if !self.is_game_over() {
            return None
        }
        println!("frames {:?}", self.frames);
        let mut total = 0;
        for (i, frame) in self.frames.iter().enumerate() {
            let (next_roll1, next_roll2) = self.get_next_two_rolls(i);
            println!("frame {} next_roll1 {:?} next_roll2 {:?}", i, next_roll1, next_roll2);
            let frame_score = frame.score(next_roll1, next_roll2);
            println!("frame {} score {}", i, frame_score);
            total += frame_score
        }
        return Some(total)
    }

    fn is_game_over(&self) -> bool {
        if self.frames.len() != 10 {
            return false
        }
        if self.frames.last().is_some_and(|x| x.is_spare()) && self.fill_ball1.is_none() {
            return false
        }
        if self.frames.last().is_some_and(|x| x.is_strike()) && self.fill_ball2.is_none() {
            return false
        }
        return true
    }

    fn get_next_two_rolls(&self, frame_index: usize) -> (Option<u16>, Option<u16>) {
        if frame_index == 9 {
            return (self.fill_ball1, self.fill_ball2)
        }
        let next_frame = self.frames.get(frame_index+1);
        if next_frame.unwrap().is_strike() {
            let next_next_frame = self.frames.get(frame_index+2);
            if next_next_frame.is_some() {
                return (Some(next_frame.unwrap().roll1), Some(next_next_frame.unwrap().roll1))
            } else {
                return (Some(next_frame.unwrap().roll1), self.fill_ball1)
            }
        }
        return (Some(next_frame.unwrap().roll1), Some(next_frame.unwrap().roll2))
    }
}
