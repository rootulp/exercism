# Pascal's Triangle

Welcome to Pascal's Triangle on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Introduction

With the weather being great, you're not looking forward to spending an hour in a classroom.
Annoyed, you enter the class room, where you notice a strangely satisfying triangle shape on the blackboard.
Whilst waiting for your math teacher to arrive, you can't help but notice some patterns in the triangle: the outer values are all ones, each subsequent row has one more value than its previous row and the triangle is symmetrical.
Weird!

Not long after you sit down, your teacher enters the room and explains that this triangle is the famous [Pascal's triangle][wikipedia].

Over the next hour, your teacher reveals some amazing things hidden in this triangle:

- It can be used to compute how many ways you can pick K elements from N values.
- It contains the Fibonacci sequence.
- If you color odd and even numbers differently, you get a beautiful pattern called the [Sierpiński triangle][wikipedia-sierpinski-triangle].

The teacher implores you and your classmates to lookup other uses, and assures you that there are lots more!
At that moment, the school bell rings.
You realize that for the past hour, you were completely absorbed in learning about Pascal's triangle.
You quickly grab your laptop from your bag and go outside, ready to enjoy both the sunshine _and_ the wonders of Pascal's triangle.

[wikipedia]: https://en.wikipedia.org/wiki/Pascal%27s_triangle
[wikipedia-sierpinski-triangle]: https://en.wikipedia.org/wiki/Sierpi%C5%84ski_triangle

## Instructions

Your task is to output the first N rows of Pascal's triangle.

[Pascal's triangle][wikipedia] is a triangular array of positive integers.

In Pascal's triangle, the number of values in a row is equal to its row number (which starts at one).
Therefore, the first row has one value, the second row has two values, and so on.

The first (topmost) row has a single value: `1`.
Subsequent rows' values are computed by adding the numbers directly to the right and left of the current position in the previous row.

If the previous row does _not_ have a value to the left or right of the current position (which only happens for the leftmost and rightmost positions), treat that position's value as zero (effectively "ignoring" it in the summation).

## Example

Let's look at the first 5 rows of Pascal's Triangle:

```text
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
```

The topmost row has one value, which is `1`.

The leftmost and rightmost values have only one preceding position to consider, which is the position to its right respectively to its left.
With the topmost value being `1`, it follows from this that all the leftmost and rightmost values are also `1`.

The other values all have two positions to consider.
For example, the fifth row's (`1 4 6 4 1`) middle value is `6`, as the values to its left and right in the preceding row are `3` and `3`:

[wikipedia]: https://en.wikipedia.org/wiki/Pascal%27s_triangle

## Source

### Created by

- @IanWhitney

### Contributed to by

- @coriolinus
- @cwhakes
- @efx
- @ErikSchierboom
- @hekrause
- @lutostag
- @navossoc
- @nfiles
- @petertseng
- @rofrol
- @stringparser
- @xakon
- @ZapAnton

### Based on

Pascal's Triangle at Wolfram Math World - https://www.wolframalpha.com/input/?i=Pascal%27s+triangle