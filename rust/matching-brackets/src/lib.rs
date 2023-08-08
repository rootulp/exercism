pub fn brackets_are_balanced(input: &str) -> bool {
    if input.is_empty() {
        return true;
    }
    let mut stack = Vec::new();
    for c in input.chars() {
        if is_opening_bracket(c) {
            stack.push(c);
        }
        if is_closing_bracket(c) {
            let opening_bracket = get_opening_bracket(c);
            match stack.pop() {
                Some(opening) => {
                    if opening != opening_bracket {
                        return false;
                    }
                }
                None => return false,
            }
        }
    }
    stack.is_empty()
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
