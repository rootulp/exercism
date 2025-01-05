# Series

Welcome to Series on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Given a string of digits, output all the contiguous substrings of length `n` in that string in the order that they appear.

For example, the string "49142" has the following 3-digit series:

- "491"
- "914"
- "142"

And the following 4-digit series:

- "4914"
- "9142"

And if you ask for a 6-digit series from a 5-digit string, you deserve whatever you get.

Note that these series are only required to occupy _adjacent positions_ in the input;
the digits need not be _numerically consecutive_.

Different languages on Exercism have different expectations about what the result should be if the length of the substrings is zero.
On the Rust track, we don't have a test for that case, so you are free to do what you feel is most appropriate.

Consider the advantages and disadvantages of the following possibilities:
- Crash the program with `panic!`.
- Return a `Result::Err`. (not possible here, because the function signature is given)
- Return an empty vector.
- Return a vector containing as many empty strings as the length of the string "digits" **plus one**. (this has some nice mathematical properties!)

## Source

### Created by

- @kedeggel

### Contributed to by

- @Baelyk
- @coriolinus
- @cwhakes
- @efx
- @ErikSchierboom
- @lutostag
- @petertseng
- @rofrol
- @stringparser
- @xakon
- @ZapAnton

### Based on

A subset of the Problem 8 at Project Euler - https://projecteuler.net/problem=8