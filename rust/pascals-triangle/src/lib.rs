pub struct PascalsTriangle {
    rows: Vec<Vec<u32>>,
}

impl PascalsTriangle {
    pub fn new(row_count: u32) -> Self {
        let mut rows = vec![];
        for i in 1..=row_count {
            let row = generate_row(i);
            rows.push(row)
        }

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
    let mut row = vec![1];
    let prev_row = generate_row(row_index - 1);
    for window in prev_row.windows(2) {
        row.push(window[0] + window[1])
    }
    row.push(1);
    row
}
