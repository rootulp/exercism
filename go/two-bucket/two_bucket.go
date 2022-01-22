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
	fmt.Printf("Cloning bucket %v with capacity %v and level %v\n", b.name, b.capacity, level)
	return NewBucket(b.name, b.capacity, level)
}
func (b *Bucket) Fill() *Bucket {
	return b.Clone(b.capacity)
}
func (b *Bucket) Empty() *Bucket {
	return b.Clone(0)
}
func (b *Bucket) Transfer(destination *Bucket) (resultSrc *Bucket, resultDestination *Bucket) {
	if b.level+destination.level > destination.capacity {
		return b.Clone(destination.capacity - destination.level), destination.Clone(destination.capacity)
	} else {
		fmt.Printf("destination.level %v, b.level %v\n", destination.level, b.level)
		return b.Clone(0), destination.Clone(destination.level + b.level)
	}
}

type state struct {
	a     *Bucket
	b     *Bucket
	moves int
}

func (s state) String() string {
	return fmt.Sprintf("a %v, b %v, moves %v\n", s.a, s.b, s.moves)
}

func Solve(bucketOneCapacity, bucketTwoCapacity, goalAmount int, startBucket string) (goalBucket string, moves int, otherBucketLevel int, e error) {

	one := NewBucket("one", bucketOneCapacity, 0)
	two := NewBucket("two", bucketTwoCapacity, 0)
	moves = 0
	queue := []state{}
	seen := map[state]bool{}

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
		fmt.Printf("current item %v, queue length: %v\n", item, len(queue))

		if _, ok := seen[item]; ok {
			continue
		} else {
			seen[item] = true
		}
		if item.a.level == goalAmount {
			return "one", item.moves, item.b.level, nil
		} else if item.b.level == goalAmount {
			return "two", item.moves, item.a.level, nil
		}

		// Execute possible moves
		// queue = append(queue, state{item.a.Fill(), item.b, moves + 1})
		// queue = append(queue, state{item.a, item.b.Fill(), moves + 1})
		// queue = append(queue, state{item.a.Empty(), item.b, moves + 1})
		// queue = append(queue, state{item.a, item.b.Empty(), moves + 1})
		// a, b := item.a.Transfer(item.b)
		// queue = append(queue, state{a, b, moves + 1})
		// a, b = item.b.Transfer(item.a)
		// queue = append(queue, state{a, b, moves + 1})
		// return ""
		fmt.Printf("current item %v, queue length: %v\n", item, len(queue))
		// return "", -1, -1, fmt.Errorf("could not find a solution")
	}

	return "", -1, -1, fmt.Errorf("could not find a solution")
}
