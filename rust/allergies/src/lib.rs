pub struct Allergies {
    score: u32,
}

#[derive(Debug, PartialEq, Eq)]
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

impl Allergen {
    pub fn value(&self) -> u32 {
        match self {
            Allergen::Eggs => 1,
            Allergen::Peanuts => 2,
            Allergen::Shellfish => 4,
            Allergen::Strawberries => 8,
            Allergen::Tomatoes => 16,
            Allergen::Chocolate => 32,
            Allergen::Pollen => 64,
            Allergen::Cats => 128,
        }
    }
}

impl Allergies {
    pub fn new(score: u32) -> Self {
        let score = score % 256;
        Self { score }
    }

    pub fn is_allergic_to(&self, allergen: &Allergen) -> bool {
        allergen.value() & self.score > 0
    }

    pub fn allergies(&self) -> Vec<Allergen> {
        let mut score = self.score;
        let mut allergies: Vec<Allergen> = Vec::new();

        let allergens = vec![
            Allergen::Cats,
            Allergen::Pollen,
            Allergen::Chocolate,
            Allergen::Tomatoes,
            Allergen::Strawberries,
            Allergen::Shellfish,
            Allergen::Peanuts,
            Allergen::Eggs,
        ];
        for candidate in allergens {
            if score / candidate.value() == 1 {
                score %= candidate.value();
                allergies.push(candidate);
            }
        }

        allergies
    }
}
