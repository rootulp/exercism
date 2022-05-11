# RPN Calculator

Welcome to RPN Calculator on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.
If you get stuck on the exercise, check out `HINTS.md`, but try and solve it without using those first :)

## Introduction

[Stacks](https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29) are a type of collection commonly used in computer science.
They are defined by their two key operations: **push** and **pop**.
**Push** adds an element to the top of the stack.
**Pop** removes and returns the topmost element.

Think of a stack like a stack of plates.
You can either add a plate to the top of the stack or take the topmost plate.
To access something further down, you have to remove all of the plates above it.

Rust's vector implementation, [`std::vec::Vec`](https://doc.rust-lang.org/std/vec/struct.Vec.html), can be used as a stack by using its `push()` and `pop()` methods.
Naturally, `push()` adds an element to the end of the `Vec` and `pop()` removes and returns the last element.
These operation can be very fast (O(1) in [Big O Notation](https://en.wikipedia.org/wiki/Big_O_notation)),
so they are one of the most idiomatic ways to use a `Vec`.

Stacks are useful to hold arbitrary numbers of elements in a specific order.
Because the last element inserted is the first element returned,
stacks are commonly refered to as **LIFO** (Last-In, First-Out).
This inherent ordering can be used for many things,
including tracking state when evaulating **Reverse Polish notation**.

## Instructions

## 1. Overview

[Reverse Polish notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation) (RPN) is a way of writing mathematical expressions.
Unlike in traditional infix notation, RPN operators *follow* their operands.
For example, instead of writing:

```
2 + 2
```

you would write:

```
2 2 +
```

The major benefit of Reverse Polish notation is that it is much simpler to parse than infix notation.
RPN eliminates the need for order of operations or parentheses in complex expressions.
For example:

```
(4 + 8) / (7 - 5)
```

can be written as

```
4 8 + 7 5 - /
```

In both cases, the expression evaluates to 6.

## 2. Example

Let's manually evaluate that complex expression.
As we learned in the introduction, evaluation of RPN requires a stack.
This stack is used to hold numeric values that the operators operate on.
We start our calculator with an empty stack and then evaluate each element one at a time.

First, we encounter a `4`,
so we push it onto our freshly created stack.

```
4
```

Next, we encounter an `8`.
We also push that onto the stack.

```
4 8
```

Now, we encounter a `+`.
We pop off the two topmost values (4 and 8),
add them together,
and push the sum back onto the stack.

```
12
```

We do something similar for `7`, `5`, and `-`:

```
12 7
12 7 5
12 2
```

Now we encounter a `/`.
Even though we last encountered a `-`,
there are two elements on the stack.
We pop off the two elements,
divide them,
and push the result back onto the stack.

```
6
```

Finally, since there is exactly one element on the stack,
we can say the expression evaluated to 6.

## 3. Goal

Your goal is to write a calculator to evaluate a list of inputs ordered by Reverse Polish notation and return the final element on the stack.

If there is not exactly one element in the stack at the end, return `None`.

If there is an operator with too few operands (such as the input `2 +`), return `None`.

You are given the following enum and stubbed function as a starting point.

```rust
#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    unimplemented!(
		"Given the inputs: {:?}, evaluate them as though they were a Reverse Polish notation expression",
		inputs,
	);
}
```

## Source

### Created by

- @cwhakes