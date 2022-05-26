pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let mut annotated: Vec<Vec<char>> = Vec::new();
    let mut field: Vec<Vec<char>> = Vec::new();

    for row in minefield {
        field.push(row.chars().collect());
    }

    for (y, row) in field.iter().enumerate() {
        let mut annotated_row = Vec::new();
        for (x, token) in row.iter().enumerate() {
            match token {
                '*' => annotated_row.push('*'),
                ' ' => {
                    let num = num_neighbor_mines(&field, y as u32, x as u32);
                    if num == 0 {
                        annotated_row.push(' ')
                    } else {
                        annotated_row.push(std::char::from_u32(num).expect("invalid num"))
                    }
                }
                _ => panic!("unsupported token"),
            }
        }
        annotated.push(annotated_row)
    }

    println!("{:?}", field);

    annotated.iter().map(|row| row.iter().collect()).collect()
}

fn num_neighbor_mines(field: &Vec<Vec<char>>, y: u32, x: u32) -> u32 {
    let neighbors = get_neighbors(field, y, x);
    let mut result = 0;

    for neighbor in neighbors {
        if neighbor == '*' {
            result += 1
        }
    }
    result
}

fn get_neighbors(field: &Vec<Vec<char>>, y: u32, x: u32) -> Vec<char> {
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
            if y + delta_y < 0 || y + delta_y >= field.len() as i32 {
                continue;
            }
            if x + delta_x < 0 || x + delta_x >= field[0].len() as i32 {
                continue;
            }
            let neighbor = field[(y + delta_y) as usize][(x + delta_y) as usize];
            neighbors.push(neighbor)
        }
    }

    neighbors
}
