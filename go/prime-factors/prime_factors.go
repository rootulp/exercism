package prime

func Factors(n int64) (factors []int64) {
	factors = []int64{}
	current := n
	divisor := int64(2)
	for current != 1 {
		if current%divisor == 0 {
			factors = append(factors, divisor)
			current = current / divisor
		} else {
			divisor += 1
		}
	}
	return factors
}
