# Card Tricks

Welcome to Card Tricks on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.
If you get stuck on the exercise, check out `HINTS.md`, but try and solve it without using those first :)

## Introduction

Slices in Go are similar to lists or arrays in other languages. They hold a number of elements of a specific type (or interface).

Slices in Go are based on arrays. Arrays have a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

A slice is written like `[]T` with `T` being the type of the elements in the slice:

```go
var empty []int                 // an empty slice
withData := []int{0,1,2,3,4,5}  // a slice pre-filled with some data
```

You can get/set an element at a given zero-based index using square-bracket notation:

```go
withData[1] = 5
x := withData[1] // x is now 5
```

You can create a new slice from an existing slice by getting a range of elements, once again using square-bracket notation, but specifying both a starting (inclusive) and ending (exclusive) index.
If you don't specify a starting index, it defaults to 0.
If you don't specify an ending index, it defaults to the length of the slice.

```go
newSlice := withData[2:4] // newSlice == []int{2,3}
```

## Instructions

As a magician-to-be, Elyse needs to practice some basics. She has a stack of cards that she wants to manipulate.

To make things a bit easier she only uses the cards 1 to 10.

## 1. Retrieve a card from a stack

Return the card at position `index` from the given stack.

```go
GetItem([]int{1, 2, 4, 1}, 2)
// Output: 4
```

## 2. Exchange a card in the stack

Exchange the card at position `index` with the new card provided and return the adjusted stack.
Note that this will also change the input slice which is ok.

```go
index := 2
new_card := 6
SetItem([]int{1, 2, 4, 1}, index, new_card)
// Output: []int{1, 2, 6, 1}
```

## 3. Create a stack of cards

Create a stack of given `length` and fill it with cards of the given `value`.

```go
PrefilledSlice(8, 3)
// Output: []int{8, 8, 8}
```

## 4. Remove a card from the stack

Remove the card at position `index` from the stack and return the stack.

```go
RemoveItem([]int{3, 2, 6, 4, 8}, 2)
// Output: []int{3, 2, 4, 8}
```

## Source

### Created by

- @tehsphinx