pub fn brackets_are_balanced(input: &str) -> bool {
    let mut seen = Vec::new();
    let brackets = input.chars().filter_map(Bracket::from_char);
    for bracket in brackets {
        match bracket {
            Bracket::Opening(_) => seen.push(bracket),
            Bracket::Closing(_) => {
                let wanted = bracket.pair();
                match seen.pop() {
                    Some(last_bracket) => {
                        if last_bracket != wanted {
                            return false;
                        }
                    }
                    None => return false,
                }
            }
        }
    }
    seen.is_empty()
}

#[derive(Debug, PartialEq)]
enum Bracket {
    Opening(char),
    Closing(char),
}

impl Bracket {
    fn from_char(c: char) -> Option<Bracket> {
        match c {
            '[' | '{' | '(' => Some(Bracket::Opening(c)),
            ']' | '}' | ')' => Some(Bracket::Closing(c)),
            _ => None,
        }
    }

    fn pair(&self) -> Bracket {
        match *self {
            Bracket::Opening(c) => Bracket::Closing(get_closing_bracket_char(c)),
            Bracket::Closing(c) => Bracket::Opening(get_opening_bracket_char(c)),
        }
    }
}

fn get_closing_bracket_char(c: char) -> char {
    match c {
        '[' => ']',
        '{' => '}',
        '(' => ')',
        _ => panic!("Invalid opening bracket: {}", c),
    }
}

fn get_opening_bracket_char(c: char) -> char {
    match c {
        ']' => '[',
        '}' => '{',
        ')' => '(',
        _ => panic!("Invalid closing bracket: {}", c),
    }
}
