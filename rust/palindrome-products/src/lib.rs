/// `Palindrome` is a newtype which only exists when the contained value is a palindrome number in base ten.
///
/// A struct with a single field which is used to constrain behavior like this is called a "newtype", and its use is
/// often referred to as the "newtype pattern". This is a fairly common pattern in Rust.
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Palindrome(u64);

impl Palindrome {
    /// Create a `Palindrome` only if `value` is in fact a palindrome when represented in base ten. Otherwise, `None`.
    pub fn new(value: u64) -> Option<Palindrome> {
        let digits = Palindrome::num_digits(value);
        if Self::is_palindrome(&digits) {
            Some(Palindrome(value))
        } else {
            None
        }
    }

    fn num_digits(num: u64) -> Vec<u64> {
        /*
         * Zero is a special case because
         * it is the terminating value of the unfold below,
         * but given a 0 as input, [0] is expected as output.
         * w/out this check, the output is an empty vec.
         */
        if num == 0 {
            return vec![0];
        }

        let mut x = num;
        let mut result = std::iter::from_fn(move || {
            if x == 0 {
                None
            } else {
                let current = x % 10;
                x /= 10;
                Some(current)
            }
        })
        .collect::<Vec<u64>>();

        result.reverse();
        result
    }

    fn is_palindrome<T>(v: &[T]) -> bool
    where
        T: Eq,
    {
        v.iter().eq(v.iter().rev())
    }

    /// Get the value of this palindrome.
    pub fn into_inner(self) -> u64 {
        return self.0;
    }
}

pub fn palindrome_products(min: u64, max: u64) -> Option<(Palindrome, Palindrome)> {
    unimplemented!(
        "returns the minimum and maximum number of palindromes of the products of two factors in the range {} to {}",
        min, max
    );
}
