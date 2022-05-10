pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        if self.health != 0_u32 {
            return None;
        }
        let mana = if self.level >= 10 { 100 } else { 0 };
        return Some(Player {
            health: 100,
            mana: Some(mana),
            level: self.level,
        });
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        unimplemented!("Cast a spell of cost {}", mana_cost)
    }
}
