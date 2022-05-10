# Role-Playing Game

Welcome to Role-Playing Game on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.
If you get stuck on the exercise, check out `HINTS.md`, but try and solve it without using those first :)

## Introduction

## Null-References

If you have ever used another programming language (C/C++, Python, Java, Ruby, Lisp, etc.), it is likely that you have encountered `null` or `nil` before. 
The use of `null` or `nil` is the way that these languages indicate that a particular variable has no value. 
However, this makes accidentally using a variable that points to `null` an easy (and frequent) mistake to make. 
As you might imagine, trying to call a function that isn't there, or access a value that doesn't exist can lead to all sorts of bugs and crashes. 
The creator of `null` went so far as to call it his ['billion-dollar mistake'.](https://www.infoq.com/presentations/Null-References-The-Billion-Dollar-Mistake-Tony-Hoare/)

## The `Option` Type

To avoid these problems, Rust does not use null-references.
However, it still needs a safe way to indicate that a particular variable has no value.
This is where `Option` comes in.
Instead of having a variable which lacks a value, Rust variables can use the `Option` enum.
This enum has two variants: `None`, Rust's null-equivalent; and `Some(T)`, where T is a value of any type.

It looks like this:

```rust
enum Option<T> {
    None,
    Some(T),
}
```

You can think of `Option` as a layer of safety between you and the problems that null-references can cause, while still retaining their conceptual usefulness.

## Using `Option`

Setting a variable to `None` is fairly straightforward:

```rust
let nothing: Option<u32> = None;  // Variable nothing is set to None
```

However, if you wish for the `Option` type to carry a value, you cannot assign this value directly.
An `Option` type variable and, say, an `i32` type variable are not equivalent.
You will need to use `Some`:

```rust
let wrong_way: Option<i32> = -4; // This will not work

let right_way: Option<i32> = Some(-4); // This will work
let another_right_way = Some(-4); // Compiler infers that this is Option<i32>
```

It's also for this reason that the following will not work:

```rust
let number = 47;
let option_number = Some(15);

let compile_error = number + option_number; // Cannot add an i32 and an Option<i32> - they are of different types
```

If you wish to get the value that is contained by Some, you will first need to check that it exists:

```rust
let mut some_words = Some("choose something to say"); // some_words set to something

match some_words {
    Some(str) => println!("Here, we will {}", str),
    None => println!("I've got nothing to say"),
} // Prints "Here, we will choose something to say"

some_words = None; // some_words now set to None

// exactly the same match block as above
match some_words {
    Some(str) => println!("Here, we will {}", str),
    None => println!("I've got nothing to say"),
} // Prints "I've got nothing to say"
```

Besides `match`, Rust has other tools available for checking and accessing values contained within `Option`, but `match` should be familiar to you by now.

Additionally, consider this a demonstration of why Rust uses `Option` instead of a null-reference.
The point is that **you _must_ check** whether or not the `Option` variable is `Some` (in which case you can go ahead and extract and use the value contained within), or `None`.
Anything else, and your program will not compile; the compiler is keeping you safe from `null`.

## Instructions

You're working on implementing a role-playing game. The player's character is represented by the following:

```rust
pub struct Player {
    health: u32,
    mana: Option<u32>,
    level: u32,
}
```

Players in this game must reach level 10 before they unlock a mana pool so that they can start casting spells. Before that point, the Player's mana is `None`.

You're working on two pieces of functionality in this game, the revive mechanic and the spell casting mechanic.

The `revive` method should check to ensure that the Player is indeed dead (their health has reached 0), and if they are, the method should return a new Player instance with 100 health.
If the Player's level is 10 or above, they should also be revived with 100 mana.
If the Player's level is below 10, their mana should be `None`. The `revive` method should preserve the Player's level.

```rust
let dead_player = Player { health: 0, mana: None, level: 2 };
dead_player.revive()
// Returns Player { health: 100, mana: None, level: 2 }
```

If the `revive` method is called on a Player whose health is 1 or above, then the method should return `None`.

```rust
let alive_player = Player { health: 1, mana: Some(15), level: 11 };
alive_player.revive()
// Returns None
```

The `cast_spell` method takes a mutable reference to the Player as well as a `mana_cost` parameter indicating how much mana the spell costs. It returns the amount of damage that the cast spell performs, which will always be two times the mana cost of the spell if the spell is successfully cast.

- If the player does not have access to a mana pool, attempting to cast the spell must decrease their health by the mana cost of the spell. The damage returned must be 0.

  ```rust
  let not_a_wizard_yet = Player { health: 79, mana: None, level: 9 };
  assert_eq!(not_a_wizard_yet.cast_spell(5), 0)
  assert_eq!(not_a_wizard_yet.health, 74);
  assert_eq!(not_a_wizard_yet.mana, None);
  ```

- If the player has a mana pool but insufficient mana, the method should not affect the pool, but instead return 0

  ```rust
  let low_mana_wizard = Player { health: 93, mana: Some(3), level: 12 };
  assert_eq!(low_mana_wizard.cast_spell(10), 0);
  assert_eq!(low_mana_wizard.health, 93);
  assert_eq!(low_mana_wizard.mana, Some(3));
  ```

- Otherwise, the `mana_cost` should be deducted from the Player's mana pool and the appropriate amount of damage should be returned.

  ```rust
  let wizard = Player { health: 123, mana: Some(30), level: 18 };
  assert_eq!(wizard.cast_spell(10), 20);
  assert_eq!(wizard.health, 123);
  assert_eq!(wizard.mana, Some(20));
  ```

Have fun!

## Source

### Created by

- @seanchen1991
- @coriolinus

### Contributed to by

- @PaulT89