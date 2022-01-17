package twobucket

import "fmt"

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	tb := TwoBucket{
		sizeBucketOne: sizeBucketOne,
		sizeBucketTwo: sizeBucketTwo,
		goalAmount:    goalAmount,
		startBucket:   startBucket,
		moves:         0,
		bucketOne:     0,
		bucketTwo:     0,
	}
	return tb.Solve()
}

type TwoBucket struct {
	sizeBucketOne int
	sizeBucketTwo int
	goalAmount    int
	startBucket   string

	// state
	moves     int
	bucketOne int
	bucketTwo int
}

func (tb TwoBucket) Solve() (goalBucket string, moves int, otherBucket int, e error) {
	if tb.bucketOne == tb.goalAmount {
		return "one", tb.moves, tb.bucketTwo, nil
	} else if tb.bucketTwo == tb.goalAmount {
		return "two", tb.moves, tb.bucketOne, nil
	}

	if tb.moves == 0 {
		tb.fillBucket(tb.startBucket)
		return tb.Solve()
	}

	tb1 := tb.Clone()

}

func (tb TwoBucket) fillBucket(bucketToFill string) {
	if bucketToFill == "one" {
		tb.bucketOne = tb.sizeBucketOne
		tb.moves += 1
	} else if bucketToFill == "two" {
		tb.bucketTwo = tb.sizeBucketTwo
		tb.moves += 1
	} else {
		panic(fmt.Sprintf("invalid bucketToFill %s", bucketToFill))
	}
}

func (tb TwoBucket) Clone() (cloned TwoBucket) {
	cloned = TwoBucket{
		sizeBucketOne: tb.sizeBucketOne,
		sizeBucketTwo: tb.sizeBucketTwo,
		goalAmount:    tb.goalAmount,
		startBucket:   tb.startBucket,

		// state
		moves:     tb.moves,
		bucketOne: tb.bucketOne,
		bucketTwo: tb.bucketTwo,
	}
	return cloned
}
