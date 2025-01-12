pub struct PascalsTriangle{
    rows: Vec<Vec<u32>>,
}

impl PascalsTriangle {
    pub fn new(row_count: u32) -> Self {
        let rows = generate_rows(row_count);
        PascalsTriangle{
            rows: rows
        }
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        return self.rows.clone()
    }
}

fn generate_rows(row_count: u32) -> Vec<Vec<u32>> {
    let mut rows = vec!();
    for i in 1..=row_count {
        let row = generate_row(i);
        rows.push(row)
    }
    rows
}

fn generate_row(row_index: u32) -> Vec<u32> {
    if row_index == 1 {
        return vec!(1)
    }
    vec!()
}