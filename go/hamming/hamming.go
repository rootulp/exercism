package hamming

import "errors"

// Distance returns the hamming distance between two DNA strands A and B.
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return distance, errors.New("a and b must be of equal length")
	}

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
