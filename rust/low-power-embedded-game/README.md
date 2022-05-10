# Low-Power Embedded Game

Welcome to Low-Power Embedded Game on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.
If you get stuck on the exercise, check out `HINTS.md`, but try and solve it without using those first :)

## Introduction

Tuples are a lightweight way to group a fixed set of arbitrary types of data together. A tuple doesn't have
a particular name; naming a data structure turns it into a `struct`. A tuple's fields don't have names;
they are accessed by means of destructuring or by position.

## Syntax

### Creation

Tuples are always created with a tuple expression:

```rust
// pointless but legal
let unit = ();
// single element
let single_element = ("note the comma",);
// two element
let two_element = (123, "elements can be of differing types");
```

Tuples can have an arbitrary number of elements.

### Access by destructuring

It is possible to access the elements of a tuple by destructuring. This just means assigning variable
names to the individual elements of the tuple, consuming it.

```rust
let (elem1, _elem2) = two_element;
assert_eq!(elem1, 123);
```

### Access by position

It is also possible to access the elements of a tuple by numeric positional index. Indexing, as always,
begins at 0.

```rust
let notation = single_element.0;
assert_eq!(notation, "note the comma");
```

## Tuple Structs

You will also be asked to work with tuple structs. Like normal structs, these are named types; unlike
normal structs, they have anonymous fields. Their syntax is very similar to normal tuple syntax. It is
legal to use both destructuring and positional access.

```rust
struct TupleStruct(u8, i32);
let my_tuple_struct = TupleStruct(123, -321);
let neg = my_tuple_struct.1;
let TupleStruct(byte, _) = my_tuple_struct;
assert_eq!(neg, -321);
assert_eq!(byte, 123);
```

### Field Visibility

All fields of anonymous tuples are always public. However, fields of tuple structs have individual
visibility which defaults to private, just like fields of standard structs. You can make the fields
public with the `pub` modifier, just as in a standard struct.

```rust
// fails due to private fields
mod tuple { pub struct TupleStruct(u8, i32); }
fn main() { let _my_tuple_struct = tuple::TupleStruct(123, -321); }
```

```rust
// succeeds: fields are public
mod tuple { pub struct TupleStruct(pub u8, pub i32); }
fn main() { let _my_tuple_struct = tuple::TupleStruct(123, -321); }
```

## Instructions

You are working on a game targeting a low-power embedded system and need to write several convenience functions which will be used by other parts of the game.

## 1. Calculate the quotient and remainder of a division

A quotient is the output of a division.

```rust
fn divmod(dividend: i16, divisor: i16) -> (i16, i16)
```

Example:

```rust
assert_eq!(divmod(10, 3), (3, 1));
```

## 2. Choose even-positioned items from an iterator

This will be helpful to enable a screen-buffer optimization, your boss promises.

Iterators are items which expose the methods defined by the [`Iterator` trait](https://doc.rust-lang.org/std/iter/trait.Iterator.html). That documentation is fairly extensive, because they offer many methods; here are the most relevant properties:

- An iterator is an arbitrary-length stream of items
- They have an [`enumerate` method](https://doc.rust-lang.org/std/iter/trait.Iterator.html#method.enumerate) which returns a tuple `(i, val)` for each value
- They have a [`filter` method](https://doc.rust-lang.org/std/iter/trait.Iterator.html#method.filter) which uses a closure to determine whether to yield an element of the iterator
- They have a [`map` method](https://doc.rust-lang.org/std/iter/trait.Iterator.html#method.map) which uses a closure to modify elements of the iterator

Because your function can run on any kind of iterator, it uses `impl` to signify that this is a trait instance instead of a simple item. Likewise, the `<Item=T>` syntax just means that it doesn't matter what kind of item the iterator produces; your function can produce the even elements of any iterator.

```rust
fn evens<T>(iter: impl Iterator<Item=T>) -> impl Iterator<Item=T>
```

Examples:

```rust
let mut even_ints = evens(0_u8..);
assert_eq!(even_ints.next(), Some(0));
assert_eq!(even_ints.next(), Some(2));
assert_eq!(even_ints.next(), Some(4));
assert_eq!(even_ints.next(), Some(6));
```

```rust
let mut evens_from_odds = evens(1_i16..);
assert_eq!(evens_from_odds.next(), Some(1));
assert_eq!(evens_from_odds.next(), Some(3));
assert_eq!(evens_from_odds.next(), Some(5));
assert_eq!(evens_from_odds.next(), Some(7));
```

## 3. Calculate the manhattan distance of a position from the origin

For mapping convenience, you have a tuple struct `Position`:

```rust
struct Position(i16, i16);
```

You need to implement a method `manhattan` on `Position` which returns the [manhattan distance](https://en.wikipedia.org/wiki/Taxicab_geometry) of that position from the origin (`Position(0, 0)`).

```rust
impl Position {
    fn manhattan(&self) -> i16
}
```

Example:

```rust
assert_eq!(Position(3, 4).manhattan(), 7);
```

## Source

### Created by

- @coriolinus