package twobucket

import "fmt"

type Bucket struct {
	name     string
	capacity int
	level    int
}

func NewBucket(name string, capacity int, level int) *Bucket {
	return &Bucket{name, capacity, level}
}
func (b *Bucket) Clone(level int) *Bucket {
	return NewBucket(b.name, b.capacity, level)
}
func (b *Bucket) Fill() *Bucket {
	return b.Clone(b.capacity)
}
func (b *Bucket) Empty() *Bucket {
	return b.Clone(0)
}
func (b *Bucket) Transfer(destination *Bucket) (src *Bucket, dst *Bucket) {
	if b.level+destination.level > destination.capacity {
		return src.Clone(destination.capacity - destination.level), destination.Clone(destination.capacity)
	} else {
		return src.Clone(0), destination.Clone(destination.level + src.level)
	}
}

type state struct {
	a     *Bucket
	b     *Bucket
	moves int
}

func Solve(bucketOneCapacity, bucketTwoCapacity, goalAmount int, startBucket string) (goalBucket string, moves int, otherBucketLevel int, e error) {

	one := NewBucket("one", bucketOneCapacity, 0)
	two := NewBucket("two", bucketTwoCapacity, 0)
	moves = 0
	queue := []state{}

	switch startBucket {
	case "one":
		one.Fill()
		moves += 1
		queue = append(queue, state{one, two, moves})
	case "two":
		two.Fill()
		moves += 1
		queue = append(queue, state{one, two, moves})
	default:
		return "", -1, -1, fmt.Errorf("start bucket invalid %s", startBucket)
	}

	for len(queue) > 0 {
		item, queue := queue[0], queue[1:]
		if item.a.level == goalAmount {
			return "one", item.moves, item.b.level, nil
		} else if item.b.level == goalAmount {
			return "two", item.moves, item.a.level, nil
		}

		// Execute possible moves
		queue = append(queue, state{item.a.Fill(), item.b, moves + 1})
		queue = append(queue, state{item.a, item.b.Fill(), moves + 1})
		queue = append(queue, state{item.a.Empty(), item.b, moves + 1})
		queue = append(queue, state{item.a, item.b.Empty(), moves + 1})
		a, b := item.a.Transfer(item.b)
		queue = append(queue, state{a, b, moves + 1})
		a, b = item.b.Transfer(item.a)
		queue = append(queue, state{a, b, moves + 1})
	}

	return "", -1, -1, fmt.Errorf("could not find a solution")
}
