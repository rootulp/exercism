use std::collections::HashMap;

pub struct Allergies {
    score: u32,
}

#[derive(Debug, PartialEq)]
pub enum Allergen {
    Eggs,
    Peanuts,
    Shellfish,
    Strawberries,
    Tomatoes,
    Chocolate,
    Pollen,
    Cats,
}

impl Allergies {
    pub fn new(score: u32) -> Self {
        return Allergies { score };
    }

    pub fn is_allergic_to(&self, allergen: &Allergen) -> bool {
        for candidate in self.allergies() {
            if candidate == allergen {
                return true;
            }
        }
        return false;
    }

    pub fn allergies(&self) -> Vec<&Allergen> {
        let score_to_allergen: HashMap<u32, &Allergen> = HashMap::from([
            (1, &Allergen::Eggs),
            (2, &Allergen::Peanuts),
            (4, &Allergen::Shellfish),
            (8, &Allergen::Strawberries),
            (16, &Allergen::Tomatoes),
            (32, &Allergen::Chocolate),
            (64, &Allergen::Pollen),
            (128, &Allergen::Cats),
        ]);

        let mut score: u32 = self.score;
        let mut allergies: Vec<&Allergen> = vec![];
        if score >= 256 {
            score %= 256
        }
        for key in self.descending_keys(score_to_allergen) {
            let allergen = score_to_allergen
                .get(&key)
                .expect("expected allergy to exist");

            if score / key == 1 {
                allergies.push(allergen);
                score %= key;
            }
        }
        allergies
    }

    fn descending_keys(&self, map: HashMap<u32, &Allergen>) -> Vec<u32> {
        let mut keys: Vec<u32> = map.into_keys().collect();
        keys.sort();
        keys
    }
}
