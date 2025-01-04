#[derive(Clone, Copy, PartialEq)]
#[repr(u8)]
enum Token {
    Mine = b'*',
    Empty = b' ',
}

impl From<Token> for char {
    fn from(t: Token) -> Self {
        t as u8 as char
    }
}

impl Token {
    fn from(c: char) -> Result<Token, ()> {
        match c {
            '*' => Ok(Token::Mine),
            ' ' => Ok(Token::Empty),
            _ => Err(()),
        }
    }
}

struct Board {
    field: Vec<Vec<Token>>,
}

impl Board {
    fn new(minefield: &[&str]) -> Self {
        let field = minefield
            .iter()
            .map(|row| row.chars().map(|char| Token::from(char).unwrap()).collect())
            .collect();
        Board { field }
    }

    fn annotate(&self) -> Vec<String> {
        self.field
            .iter()
            .enumerate()
            .map(|(row, _)| self.annotate_row(row))
            .collect()
    }

    fn annotate_row(&self, y: usize) -> String {
        let row = &self.field[y];
        row.iter()
            .enumerate()
            .map(|(x, _)| self.annotate_location(y, x))
            .collect()
    }

    fn annotate_location(&self, y: usize, x: usize) -> char {
        let token = self.field[y][x];
        match token {
            Token::Mine => Token::Mine.into(),
            Token::Empty => {
                let neighbor_mines = self.num_neighbor_mines(y as u32, x as u32);
                if neighbor_mines == 0 {
                    ' '
                } else {
                    (b'0' + neighbor_mines as u8) as char
                }
            }
        }
    }

    fn num_neighbor_mines(&self, y: u32, x: u32) -> usize {
        self.get_neighbors(y, x)
            .iter()
            .filter(|neighbor| neighbor == &&&Token::Mine)
            .count()
    }

    fn get_neighbors(&self, y: u32, x: u32) -> Vec<&Token> {
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
                if y + delta_y < 0 || y + delta_y >= self.num_rows() as i32 {
                    continue;
                }
                if x + delta_x < 0 || x + delta_x >= self.num_cols() as i32 {
                    continue;
                }
                let neighbor = &self.field[(y + delta_y) as usize][(x + delta_x) as usize];
                neighbors.push(neighbor)
            }
        }

        neighbors
    }

    fn num_rows(&self) -> usize {
        self.field.len()
    }

    fn num_cols(&self) -> usize {
        self.field[0].len()
    }
}

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let board = Board::new(minefield);
    board.annotate()
}
