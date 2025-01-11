#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    NotEnoughPinsLeft,
    GameComplete,
}

pub struct BowlingGame {
    frames: Vec<Frame>,
    previous_roll: Option<u16>,
    fill_ball1: Option<u16>,
    fill_ball2: Option<u16>
}

impl BowlingGame {
    pub fn new() -> Self {
        return BowlingGame{
            frames: vec!(),
            previous_roll: None,
            fill_ball1: None,
            fill_ball2: None,
        }
    }

    pub fn roll(&mut self, pins: u16) -> Result<(), Error> {
        if self.is_game_over() {
            return Err(Error::GameComplete)
        }
        if pins > 10 {
            return Err(Error::NotEnoughPinsLeft)
        }

        // Handle fill balls
        if self.frames.last().is_some_and(|x| x.is_spare()) {
            self.fill_ball1 = Some(pins);
        }
        if self.frames.last().is_some_and(|x| x.is_strike()) {
            if self.fill_ball1.is_none() {
                self.fill_ball1 = Some(pins);
            } else {
                self.fill_ball2 = Some(pins);
            }
        }

        // Handle normal rolls
        if pins == 10 {
            let frame = Frame::new(pins, 0);
            self.frames.push(frame);
            return Ok(())
        }
        if self.previous_roll.is_some() {
            let roll1 = self.previous_roll.unwrap();
            let frame = Frame::new(roll1, pins);
            self.frames.push(frame);
            self.previous_roll = None;
        } else {
            self.previous_roll = Some(pins);
        }
        Ok(())
    }

    pub fn score(&self) -> Option<u16> {
        if !self.is_game_over() {
            return None
        }
        println!("frames {:?}", self.frames);
        let mut total = 0;
        for (i, frame) in self.frames.iter().enumerate() {
            // let next_frame = self.frames.get(i+1);
            let next_frame = self.frames.get(i+1);
            println!("frame {:?} next_frame {:?}", frame, next_frame);

            if next_frame.is_none() {
                let frame_score = frame.score(self.fill_ball1.unwrap_or(0), self.fill_ball2.unwrap_or(0));
                println!("frame {} score {}", i, frame_score);
                total += frame_score
            } else {
                let next_roll1 = next_frame.unwrap().roll1;
                let next_roll2 = next_frame.unwrap().roll2;
                println!("frame {} next_roll1 {} next_roll2 {}", i, next_roll1, next_roll2);
                let frame_score =  frame.score(next_roll1, next_roll2);
                println!("frame {} score {}", i, frame_score);
                total += frame_score
            }
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
}

#[derive(Debug, PartialEq, Eq)]
pub struct Frame {
    pub roll1: u16,
    pub roll2: u16
}

impl Frame {
    pub fn new(roll1:u16, roll2: u16) -> Self {
        return Frame{
            roll1: roll1,
            roll2: roll2
        }
    }

    pub fn score(&self, next_roll1: u16, next_roll2: u16) -> u16 {
        if self.is_strike() {
            return 10 + next_roll1 + next_roll2;
        }
        if self.is_spare() {
            return 10 + next_roll1;
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
        self.roll1 + self.roll2
    }
}
