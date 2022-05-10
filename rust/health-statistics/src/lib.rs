pub struct User {
    name: String,
    age: u32,
    weight: f32,
}

impl User {
    pub fn new(name: String, age: u32, weight: f32) -> Self {
        return User { name, age, weight };
    }

    pub fn name(&self) -> &str {
        return &self.name;
    }

    pub fn age(&self) -> u32 {
        return self.age;
    }

    pub fn weight(&self) -> f32 {
        return self.weight;
    }

    pub fn set_age(&mut self, new_age: u32) {
        self.age = new_age
    }

    pub fn set_weight(&mut self, new_weight: f32) {
        self.weight = new_weight
    }
}
