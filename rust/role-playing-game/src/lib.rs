use std::cmp;

pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        if self.health > 0 {
            None
        } else {
            Some(Player {
                health: 100,
                mana: if self.level >= 10 { Some(100) } else { None },
                level: self.level,
            })
        }
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match self.mana {
            None => {
                self.health -= cmp::min(mana_cost, self.health);
                0
            }
            Some(mana) if mana < mana_cost => 0,
            Some(mana) => {
                self.mana = Some(mana - mana_cost);
                mana_cost * 2
            }
        }
    }
}
