pub fn find(array: &[i32], key: i32) -> Option<usize> {
    for (index, element) in array.iter().enumerate() {
        if element == &key {
            return Some(index)
        }
    }
    None
}
