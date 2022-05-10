pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        if self.health != 0_u32 {
            return None;
        } else if self.level >= 10 {
            return Some(Player {
                health: 100,
                mana: Some(100),
                level: self.level,
            });
        } else {
            return Some(Player {
                health: 100,
                mana: Some(0),
                level: self.level,
            });
        }
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        unimplemented!("Cast a spell of cost {}", mana_cost)
    }
}
