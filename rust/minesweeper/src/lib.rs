struct Board {
    field: Vec<Vec<char>>,
}

impl Board {
    fn new(minefield: &[&str]) -> Self {
        let mut field: Vec<Vec<char>> = Vec::new();

        for row in minefield {
            field.push(row.chars().collect());
        }
        Board { field }
    }

    fn annotated(&self) -> Vec<String> {
        let mut annotated: Vec<String> = Vec::new();

        for (y, row) in self.field.iter().enumerate() {
            let mut annotated_row = String::new();
            for (x, token) in row.iter().enumerate() {
                match token {
                    '*' => annotated_row.push('*'),
                    ' ' => {
                        let neighbor_mines = self.num_neighbor_mines(y as u32, x as u32);
                        let annotated_location = if neighbor_mines == 0 {
                            " ".to_string()
                        } else {
                            neighbor_mines.to_string()
                        };
                        annotated_row.push_str(&annotated_location)
                    }
                    _ => panic!("unsupported token"),
                }
            }
            annotated.push(annotated_row)
        }

        annotated
    }

    fn num_neighbor_mines(&self, y: u32, x: u32) -> usize {
        return self
            .get_neighbors(y, x)
            .iter()
            .filter(|neighbor| neighbor == &&&'*')
            .count();
    }

    fn get_neighbors(&self, y: u32, x: u32) -> Vec<&char> {
        let mut neighbors = Vec::new();
        let y = y as i32;
        let x = x as i32;

        let delta_ys: Vec<i32> = (-1..=1).collect();
        for delta_y in delta_ys {
            let delta_xs: Vec<i32> = (-1..=1).collect();
            for delta_x in delta_xs {
                if delta_y == 0 && delta_x == 0 {
                    continue;
                }
                if y + delta_y < 0 || y + delta_y >= self.field.len() as i32 {
                    continue;
                }
                if x + delta_x < 0 || x + delta_x >= self.field[0].len() as i32 {
                    continue;
                }
                let neighbor = &self.field[(y + delta_y) as usize][(x + delta_x) as usize];
                neighbors.push(neighbor)
            }
        }

        neighbors
    }
}

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let board = Board::new(minefield);
    return board.annotated();
}
