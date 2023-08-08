pub fn brackets_are_balanced(input: &str) -> bool {
    let mut seen = Vec::new();
    let brackets = input.chars().filter_map(Bracket::from_char);
    for bracket in brackets {
        match bracket {
            Bracket::Opening(c) => seen.push(c),
            Bracket::Closing(c) => {
                let wanted_bracket = get_opening_bracket(c);
                match seen.pop() {
                    Some(last_bracket) => {
                        if last_bracket != wanted_bracket {
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
}

fn get_opening_bracket(c: char) -> char {
    match c {
        ']' => '[',
        '}' => '{',
        ')' => '(',
        _ => panic!("Invalid closing bracket"),
    }
}
