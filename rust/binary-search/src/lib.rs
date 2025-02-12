pub fn find(array: &[i32], key: i32) -> Option<usize> {
    if array.is_empty() {
        return None
    }

    let mut left = 0;
    let mut right = array.len() - 1;
    let mut middle;

    while left <= right {
        middle = get_middle(left, right);
        println!("left: {}, right: {}, middle: {}", left, right, middle);

        // Special case
        if middle == 0 {
            if array[middle] == key {
                return Some(middle)
            } else {
                return None
            }
        }


        if array[middle] == key {
            return Some(middle)
        } else if array[middle] < key {
            left = middle + 1;
        } else {
            right = middle - 1;
        }
    }
    None
}

fn get_middle(left: usize, right :usize) -> usize {
    left + ((right - left) / 2)
}
