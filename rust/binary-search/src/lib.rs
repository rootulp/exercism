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

        if out_of_bounds(left, right, middle, array) {
            return None
        }

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

fn out_of_bounds(left: usize, right: usize, middle: usize, array: &[i32]) -> bool {
    if left > array.len() {
        return true
    }
    if right > array.len() {
        return true
    }
    if middle > array.len() {
        return true
    }
    false
}
