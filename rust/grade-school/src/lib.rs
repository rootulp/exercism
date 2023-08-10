use std::collections::HashMap;

pub struct School {
    grade_to_names: HashMap<u32, Vec<String>>,
}

impl School {
    pub fn new() -> School {
        School {
            grade_to_names: HashMap::new(),
        }
    }

    pub fn add(&mut self, grade: u32, student: &str) {
        if let Some(existing) = self.grade_to_names.get_mut(&grade) {
            existing.push(student.to_string());
            existing.sort()
        } else {
            self.grade_to_names.insert(grade, vec![student.to_string()]);
        }
    }

    pub fn grades(&self) -> Vec<u32> {
        let mut keys: Vec<u32> = self.grade_to_names.keys().cloned().collect();
        keys.sort();
        keys
    }

    // If `grade` returned a reference, `School` would be forced to keep a `Vec<String>`
    // internally to lend out. By returning an owned vector of owned `String`s instead,
    // the internal structure can be completely arbitrary. The tradeoff is that some data
    // must be copied each time `grade` is called.
    pub fn grade(&self, grade: u32) -> Vec<String> {
        self.grade_to_names.get(&grade).cloned().unwrap_or_default()
    }
}

impl Default for School {
    fn default() -> Self {
        Self::new()
    }
}
