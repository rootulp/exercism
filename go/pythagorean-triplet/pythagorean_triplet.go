package pythagorean

import "math"

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (result []Triplet) {
	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
			if c == math.Trunc(c) && c <= float64(max) && float64(b) < c {
				result = append(result, Triplet{a, b, int(c)})
			}
		}
	}
	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.
func Sum(p int) (result []Triplet) {
	candidates := Range(0, p)
	for _, candidate := range candidates {
		if candidate[0]+candidate[1]+candidate[2] == p {
			result = append(result, candidate)
		}
	}
	return result
}

func isTriplet(a int, b int, c int) bool {
	return math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2)
}
