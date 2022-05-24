# Clock

Welcome to Clock on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Implement a clock that handles times without dates.

You should be able to add and subtract minutes to it.

Two clocks that represent the same time should be equal to each other.

You will also need to implement `.to_string()` for the `Clock` struct. We will be using this to display the Clock's state.  You can either do it via implementing it directly or using the [Display trait](https://doc.rust-lang.org/std/fmt/trait.Display.html).

Did you implement `.to_string()` for the `Clock` struct?

If so, try implementing the
[Display trait](https://doc.rust-lang.org/std/fmt/trait.Display.html) for `Clock` instead.

Traits allow for a common way to implement functionality for various types.

For additional learning, consider how you might implement `String::from` for the `Clock` type.
You don't have to actually implement this—it's redundant with `Display`, which is generally the
better choice when the destination type is `String`—but it's useful to have a few type-conversion
traits in your toolkit.

## Source

### Created by

- @sacherjj

### Contributed to by

- @attilahorvath
- @coriolinus
- @cwhakes
- @danieljl
- @eddyp
- @efx
- @ErikSchierboom
- @felix91gr
- @kunaltyagi
- @lutostag
- @nfiles
- @petertseng
- @rofrol
- @shaaraddalvi
- @stringparser
- @tmccombs
- @xakon
- @ZapAnton

### Based on

Pairing session with Erin Drummond - https://twitter.com/ebdrummond