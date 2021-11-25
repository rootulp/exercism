package pythagorean

import "math"

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) (result []Triplet) {
	for i := min; i <= max-2; i++ {
		for j := i + 1; j <= max-1; j++ {
			for k := j + 1; k <= max; k++ {
				if isTriplet(i, j, k) {
					result = append(result, Triplet{i, j, k})
				}
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
