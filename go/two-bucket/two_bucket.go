package twobucket

import (
	"errors"
)

type Bucket struct {
	name     string
	capacity int
	level    int
}

func (b *Bucket) IsEmpty() bool {
	return b.level == 0
}
func (b *Bucket) IsFull() bool {
	return b.level == b.capacity
}
func (b *Bucket) Fill() {
	b.level = b.capacity
}
func (b *Bucket) Empty() {
	b.level = 0
}
func (src *Bucket) Transfer(dst *Bucket) {
	for src.level != 0 && dst.level != dst.capacity {
		src.level -= 1
		dst.level += 1
	}
}

func Solve(bucketOneCapacity, bucketTwoCapacity, goalAmount int, startBucket string) (goalBucket string, moves int, otherBucketLevel int, e error) {
	if !isValidCapacity(bucketOneCapacity) ||
		!isValidCapacity(bucketTwoCapacity) ||
		!isValidStartBucket(startBucket) ||
		goalAmount < 1 {
		return "", 0, 0, errors.New("invalid parameters")
	}

	one := &Bucket{"one", bucketOneCapacity, 0}
	two := &Bucket{"two", bucketTwoCapacity, 0}
	var first *Bucket
	var second *Bucket

	if startBucket == "one" {
		first, second = one, two
	} else {
		first, second = two, one
	}

	moves = 0
	for first.level != goalAmount && second.level != goalAmount {
		switch {
		case first.IsEmpty():
			first.Fill()
		case second.IsFull():
			second.Empty()
		case first.capacity == goalAmount:
			first.Fill()
		case second.capacity == goalAmount:
			second.Fill()
		default:
			first.Transfer(second)
		}
		moves += 1
	}
	if first.level == goalAmount {
		return first.name, moves, second.level, nil
	} else {
		return second.name, moves, first.level, nil
	}
}

func isValidCapacity(capacity int) bool {
	return capacity >= 1
}

func isValidStartBucket(startBucket string) bool {
	return startBucket == "one" || startBucket == "two"
}
