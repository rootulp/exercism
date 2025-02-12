pub fn find(array: &[i32], key: i32) -> Option<usize> {
    let mut left = 0;
    let mut right = array.len() - 1;
    let mut middle = get_middle(left, right);
    println!("left: {}, right: {}, middle: {}", left, right, middle);

    while left <= right {
        println!("left: {}, right: {}, middle: {}", left, right, middle);
        if array[middle] == key {
            return Some(middle)
        } else if array[middle] < key {
            left = middle + 1;
            middle = get_middle(left, right);
        } else {
            right = middle - 1;
            middle = get_middle(left, right);
        }
    }
    None
}

fn get_middle(left: usize, right :usize) -> usize {
    (right - left) / 2
}
