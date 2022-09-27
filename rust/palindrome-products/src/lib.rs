/// `Palindrome` is a newtype which only exists when the contained value is a palindrome number in base ten.
///
/// A struct with a single field which is used to constrain behavior like this is called a "newtype", and its use is
/// often referred to as the "newtype pattern". This is a fairly common pattern in Rust.
#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord)]
pub struct Palindrome(u64);

impl Palindrome {
    /// Create a `Palindrome` only if `value` is in fact a palindrome when represented in base ten. Otherwise, `None`.
    pub fn new(value: u64) -> Option<Palindrome> {
        let digits = Palindrome::digits(value);
        if Self::is_palindrome(&digits) {
            Some(Palindrome(value))
        } else {
            None
        }
    }

    fn digits(num: u64) -> Vec<u64> {
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
        self.0
    }
}

pub fn palindrome_products(min: u64, max: u64) -> Option<(Palindrome, Palindrome)> {
    let mut products: Vec<Palindrome> = vec![];
    for a in min..=max {
        for b in min..=max {
            let product = a * b;
            if let Some(palindrome) = Palindrome::new(product) {
                products.push(palindrome);
            }
        }
    }

    let min_palindrome = products.iter().min();
    let max_palindrome = products.iter().max();
    if let Some(min_palindrome) = &min_palindrome {
        if let Some(max_palindrome) = &max_palindrome {
            return Some((**min_palindrome, **max_palindrome));
        }
    }
    None
}
