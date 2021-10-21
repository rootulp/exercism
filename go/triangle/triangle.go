package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	NaT Kind = "Not a triangle"
	Equ Kind = "Equilateral"
	Iso Kind = "Isosceles"
	Sca Kind = "Scalene"
)

func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func isTriangle(a, b, c float64) bool {
	if !isValidSideLength(a) ||
		!isValidSideLength(b) ||
		!isValidSideLength(c) {
		return false
	}
	if !isSumOfTwoSidesGreaterThanThirdSide(a, b, c) ||
		!isSumOfTwoSidesGreaterThanThirdSide(a, c, b) ||
		!isSumOfTwoSidesGreaterThanThirdSide(c, b, a) {
		return false
	}
	return true
}

func isSumOfTwoSidesGreaterThanThirdSide(x, y, z float64) bool {
	return x+y >= z
}

func isValidSideLength(x float64) bool {
	return x > 0 && !math.IsNaN(x) && !math.IsInf(x, 1)
}
