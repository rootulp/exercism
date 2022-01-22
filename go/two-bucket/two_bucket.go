package twobucket

import (
	"errors"
	"fmt"
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

	switch startBucket {
	case "one":
		first, second = one, two
	case "two":
		first, second = two, one
	default:
		return "", -1, -1, fmt.Errorf("start bucket invalid %s", startBucket)
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
	} else if second.level == goalAmount {
		return second.name, moves, first.level, nil
	} else {
		return "", -1, -1, fmt.Errorf("could not find a solution")
	}
}

func isValidCapacity(capacity int) bool {
	return capacity >= 1
}

func isValidStartBucket(startBucket string) bool {
	return startBucket == "one" || startBucket == "two"
}
