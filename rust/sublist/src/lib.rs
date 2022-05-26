#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(a: &[T], b: &[T]) -> Comparison {
    if a == b {
        return Comparison::Equal;
    }
    if is_sublist(a, b) {
        return Comparison::Sublist;
    }
    if is_sublist(b, a) {
        return Comparison::Superlist;
    }
    return Comparison::Unequal;
}

fn is_sublist<T: PartialEq>(a: &[T], b: &[T]) -> bool {
    if a.len() == 0 {
        return true;
    }
    for list in b.windows(a.len()) {
        if a == list {
            return true;
        }
    }
    return false;
}
