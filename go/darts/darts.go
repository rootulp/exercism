package darts

import "math"

func Score(x, y float64) int {
	if isInnerCircle(x, y) {
		return 10
	} else if isMiddleCircle(x, y) {
		return 5
	} else if isOuterCircle(x, y) {
		return 1
	}
	return 0
}

func isInnerCircle(x float64, y float64) bool {
	return distanceToCenter(x, y) <= 1
}

func isMiddleCircle(x float64, y float64) bool {
	return distanceToCenter(x, y) <= 5
}

func isOuterCircle(x float64, y float64) bool {
	return distanceToCenter(x, y) <= 10
}

func distanceToCenter(x float64, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
