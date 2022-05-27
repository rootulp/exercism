#[derive(Clone, Copy, PartialEq)]
#[repr(u8)]
enum Token {
    Mine = b'*',
    Empty = b' ',
}

impl Into<char> for Token {
    fn into(self) -> char {
        self as u8 as char
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
        let annotated: Vec<String> = self
            .field
            .iter()
            .enumerate()
            .map(|(row, _)| self.annotate_row(row))
            .collect();
        annotated
    }

    fn annotate_row(&self, y: usize) -> String {
        let row = &self.field[y];
        let annotated_row = row
            .iter()
            .enumerate()
            .map(|(x, token)| match token {
                Token::Mine => Token::Mine.into(),
                Token::Empty => {
                    let neighbor_mines = self.num_neighbor_mines(y as u32, x as u32);
                    if neighbor_mines == 0 {
                        return ' ';
                    } else {
                        return (b'0' + neighbor_mines as u8) as char;
                    };
                }
            })
            .collect();

        annotated_row
    }

    fn num_neighbor_mines(&self, y: u32, x: u32) -> usize {
        return self
            .get_neighbors(y, x)
            .iter()
            .filter(|neighbor| neighbor == &&&Token::Mine)
            .count();
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
        return self.field.len();
    }

    fn num_cols(&self) -> usize {
        return self.field[0].len();
    }
}

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let board = Board::new(minefield);
    return board.annotate();
}
