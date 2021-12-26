package darts

import "math"

const innerRadius = 1
const middleRadius = 5
const outerRadius = 10

func Score(x, y float64) int {
	if isInCircle(x, y, innerRadius) {
		return 10
	} else if isInCircle(x, y, middleRadius) {
		return 5
	} else if isInCircle(x, y, outerRadius) {
		return 1
	}
	return 0
}

func isInCircle(x float64, y float64, radius int) bool {
	return distanceToCenter(x, y) <= float64(radius)
}

func distanceToCenter(x float64, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
