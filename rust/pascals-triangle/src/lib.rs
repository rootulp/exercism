pub struct PascalsTriangle{
    rows: Vec<Vec<u32>>,
}

impl PascalsTriangle {
    pub fn new(row_count: u32) -> Self {
        PascalsTriangle{
            rows: vec!()
        }
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        return self.rows.clone()
    }
}
