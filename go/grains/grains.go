package grains

import (
	"errors"
	"math"
)

// Square returns the number of grains on the provided square.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("square: n is not between 1 and 64 (inclusive)")
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

// Total returns the sum of all grains on the chess board.
func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i <= 64; i++ {
		var result, _ = Square(i)
		sum += result
	}
	return sum
}
