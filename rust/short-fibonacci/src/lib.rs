/// Create an empty vector
pub fn create_empty() -> Vec<u8> {
    Vec::new()
}

/// Create a buffer of `count` zeroes.
///
/// Applications often use buffers when serializing data to send over the network.
pub fn create_buffer(count: usize) -> Vec<u8> {
    let mut buffer = Vec::new();
    for _ in 0..count {
        buffer.push(0)
    }
    println!("buffer {:?}", buffer);
    buffer
}

/// Create a vector containing the first five elements of the Fibonacci sequence.
///
/// Fibonacci's sequence is the list of numbers where the next number is a sum of the previous two.
/// Its first five elements are `1, 1, 2, 3, 5`.
pub fn fibonacci() -> Vec<u8> {
    let mut result = Vec::new();
    for i in 0..5 {
        if i == 0 || i == 1 {
            result.push(1)
        } else {
            let second_to_last = result[result.len().checked_sub(2).unwrap()];
            result.push(result.last().unwrap() + second_to_last);
        }
    }
    result
}
