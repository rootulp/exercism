pub struct PascalsTriangle {
    rows: Vec<Vec<u32>>,
}

impl PascalsTriangle {
    pub fn new(row_count: u32) -> Self {
        let rows = (1..=row_count).map(generate_row).collect();
        PascalsTriangle { rows }
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        self.rows.clone()
    }
}

fn generate_row(row_index: u32) -> Vec<u32> {
    if row_index == 1 {
        return vec![1];
    }
    let prev_row = generate_row(row_index - 1);
    let mut row: Vec<u32> = prev_row
        .windows(2)
        .map(|window| window[0] + window[1])
        .collect();
    row.insert(0, 1);
    row.push(1);
    row
}
