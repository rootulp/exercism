# Two Bucket

Welcome to Two Bucket on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Given two buckets of different size and which bucket to fill first, determine how many actions are required to measure an exact number of liters by strategically transferring fluid between the buckets.

There are some rules that your solution must follow:

- You can only do one action at a time.
- There are only 3 possible actions:
  1. Pouring one bucket into the other bucket until either:
     a) the first bucket is empty
     b) the second bucket is full
  2. Emptying a bucket and doing nothing to the other.
  3. Filling a bucket and doing nothing to the other.
- After an action, you may not arrive at a state where the starting bucket is empty and the other bucket is full.

Your program will take as input:

- the size of bucket one
- the size of bucket two
- the desired number of liters to reach
- which bucket to fill first, either bucket one or bucket two

Your program should determine:

- the total number of actions it should take to reach the desired number of liters, including the first fill of the starting bucket
- which bucket should end up with the desired number of liters - either bucket one or bucket two
- how many liters are left in the other bucket

Note: any time a change is made to either or both buckets counts as one (1) action.

Example:
Bucket one can hold up to 7 liters, and bucket two can hold up to 11 liters. Let's say at a given step, bucket one is holding 7 liters and bucket two is holding 8 liters (7,8). If you empty bucket one and make no change to bucket two, leaving you with 0 liters and 8 liters respectively (0,8), that counts as one action. Instead, if you had poured from bucket one into bucket two until bucket two was full, resulting in 4 liters in bucket one and 11 liters in bucket two (4,11), that would also only count as one action.

Another Example:
Bucket one can hold 3 liters, and bucket two can hold up to 5 liters.  You are told you must start with bucket one.  So your first action is to fill bucket one.  You choose to empty bucket one for your second action.  For your third action, you may not fill bucket two, because this violates the third rule -- you may not end up in a state after any action where the starting bucket is empty and the other bucket is full.

Written with <3 at [Fullstack Academy](http://www.fullstackacademy.com/) by Lindsay Levine.

In package twobucket, implement a Go func, Solve, with
the following signature:

```
func Solve(sizeBucketOne,
	sizeBucketTwo,
	goalAmount int,
	startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error)
```
Solve returns four values: the resulting goal bucket("one" or two"),
the number of moves/steps to achieve the goal amount,
the liters left in the other bucket, and an error value.
The returned error value should be nil when the parameters are valid.
Return an error for any invalid parameter.
Solve should also return an error when a solution cannot be found,
but this error relates only to the bonus exercise below, so you may
ignore that error case for your initial solution.

## Bonus exercise

Once you get `go test` passing, try `go test -tags bonus`.  This uses a *build
tag* to enable tests that were not previously enabled. Build tags control which
files should be included in the package. You can read more about those at [the
Go documentation](https://golang.org/pkg/go/build/#hdr-Build_Constraints).

The exercise limits `go test` to only build the tests referenced in the
`two_bucket_test.go` file. The bonus test cases are found in the file
`bonus_test.go`. Enable those tests as described above.

To get the bonus tests to pass, the Solve func must detect when
a solution cannot be found and return an error.
A solution cannot be found when input test case bucket sizes
are not ones which allow the three operations to succeed in creating the goal amount,
which occurs when the two bucket sizes are not relatively prime to one another.

## Source

### Created by

- @leenipper

### Contributed to by

- @bitfield
- @ekingery
- @ferhatelmas
- @hilary
- @sebito91
- @tleen

### Based on

Water Pouring Problem - http://demonstrations.wolfram.com/WaterPouringProblem/