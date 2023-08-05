#[derive(Debug)]
pub struct HighScores<'a> {
    scores: &'a [u32],
}

impl<'a> HighScores<'a> {
    pub fn new(scores: &'a [u32]) -> Self {
        HighScores { scores }
    }

    pub fn scores(&self) -> &[u32] {
        self.scores
    }

    pub fn latest(&self) -> Option<u32> {
        return self.scores.last().copied();
    }

    pub fn personal_best(&self) -> Option<u32> {
        return self.scores.iter().max().copied();
    }

    pub fn personal_top_three(&self) -> Vec<u32> {
        let mut result = Vec::new();
        self.scores.iter().for_each(|&x| result.push(x));
        result.sort();
        result.reverse();
        result.truncate(3);
        result
    }
}
