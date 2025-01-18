pub struct Allergies {
    allergens: Vec<Allergen>,
}

#[derive(Debug, PartialEq, Eq, Clone)]
pub enum Allergen {
    Eggs,         // 1
    Peanuts,      // 2
    Shellfish,    // 4
    Strawberries, // 8
    Tomatoes,     // 16
    Chocolate,    // 32
    Pollen,       // 64
    Cats,         // 128
}

impl Allergies {
    pub fn new(score: u32) -> Self {
        let allergens = [
            Allergen::Eggs,
            Allergen::Peanuts,
            Allergen::Shellfish,
            Allergen::Strawberries,
            Allergen::Tomatoes,
            Allergen::Chocolate,
            Allergen::Pollen,
            Allergen::Cats,
        ]
        .iter()
        .enumerate()
        .filter_map(|(i, allergen)| {
            if score & (1 << i) != 0 {
                Some(allergen.clone())
            } else {
                None
            }
        })
        .collect();

        Self { allergens }
    }

    pub fn is_allergic_to(&self, allergen: &Allergen) -> bool {
        self.allergens.contains(allergen)
    }

    pub fn allergies(&self) -> Vec<Allergen> {
        self.allergens.clone()
    }
}
