use std::collections::HashMap;

pub struct School<'a> {
    grade_to_names: HashMap<u32, Vec<&'a str>>,
}

impl<'a> School<'a> {
    pub fn new() -> School<'a> {
        School {
            grade_to_names: HashMap::new(),
        }
    }

    pub fn add(&mut self, grade: u32, student: &'a str) {
        if let Some(existing) = self.grade_to_names.get_mut(&grade) {
            existing.push(student);
            existing.sort()
        } else {
            self.grade_to_names.insert(grade, vec![student]);
        }
    }

    pub fn grades(&self) -> Vec<u32> {
        let mut keys: Vec<u32> = self.grade_to_names.keys().cloned().collect();
        keys.sort();
        keys
    }

    pub fn grade(&self, grade: u32) -> Vec<String> {
        self.grade_to_names
            .get(&grade)
            .cloned()
            .unwrap_or_default()
            .into_iter()
            .map(|s| s.to_string())
            .collect()
    }
}

impl<'a> Default for School<'a> {
    fn default() -> Self {
        Self::new()
    }
}
