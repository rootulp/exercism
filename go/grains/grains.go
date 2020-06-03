package grains

import "math"

// Square returns the number of grains on the provided square.
func Square(n int) (uint64, error) {
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
