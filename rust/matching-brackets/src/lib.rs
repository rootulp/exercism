pub fn brackets_are_balanced(input: &str) -> bool {
    let mut seen = Vec::new();
    for c in input.chars() {
        if is_opening_bracket(c) {
            seen.push(c);
        }
        if is_closing_bracket(c) {
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
    seen.is_empty()
}

fn is_opening_bracket(c: char) -> bool {
    c == '[' || c == '{' || c == '('
}

fn is_closing_bracket(c: char) -> bool {
    c == ']' || c == '}' || c == ')'
}

fn get_opening_bracket(c: char) -> char {
    match c {
        ']' => '[',
        '}' => '{',
        ')' => '(',
        _ => panic!("Invalid closing bracket"),
    }
}
