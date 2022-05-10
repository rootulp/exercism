# A Short Fibonacci Sequence

Welcome to A Short Fibonacci Sequence on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.
If you get stuck on the exercise, check out `HINTS.md`, but try and solve it without using those first :)

## Introduction

Rust provides a macro `vec![]` to help you create Vectors.
This comes in quite handy when you need to initialize lists.

## Instructions

You are going to initialize empty buffers and list the first five numbers, or elements, of the Fibonacci sequence.

The Fibonacci sequence is a set of numbers where the next element is the sum of the prior two. We start the sequence at one. So the first two elements are 1 and 1.

## 1. Create a buffer of `count` zeroes.

Create a function that creates a buffer of `count` zeroes.
```rust
let my_buffer = create_buffer(5);
// [0, 0, 0, 0, 0]
```

## 2. List the first five elements of the Fibonacci sequence

Create a function that returns the first five numbers of the Fibonacci sequence. 
Its first five elements are `1, 1, 2, 3, 5`
```rust
let first_five = fibonacci();
// [1, 1, 2, 3, 5]
```

## Source

### Created by

- @efx
- @coriolinus