pub fn brackets_are_balanced(input: &str) -> bool {
    let mut seen = Vec::new();
    for b in input.chars().filter_map(Bracket::from_char) {
        match b {
            Bracket::Opening(_) => seen.push(b),
            Bracket::Closing(_) => match seen.pop() {
                Some(last) => {
                    if last != b.pair() {
                        return false;
                    }
                }
                None => return false,
            },
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
        match self {
            Self::Opening(c) => Bracket::Closing(get_pair(*c)),
            Self::Closing(c) => Bracket::Opening(get_pair(*c)),
        }
    }
}

fn get_pair(c: char) -> char {
    match c {
        '[' => ']',
        '{' => '}',
        '(' => ')',
        ']' => '[',
        '}' => '{',
        ')' => '(',
        _ => panic!("Invalid bracket: {}", c),
    }
}
