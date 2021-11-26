package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	multiples := getMultiples(limit, divisors)
	for multiple := range multiples {
		sum += multiple
	}
	return sum
}

func getMultiples(limit int, divisors []int) (set map[int]bool) {
	set = map[int]bool{}
	for _, divisor := range divisors {
		multiples := getMultiplesForDivisor(limit, divisor)
		for _, multiple := range multiples {
			set[multiple] = true
		}
	}
	return set
}

func getMultiplesForDivisor(limit int, divisor int) (multiples []int) {
	if divisor == 0 {
		return multiples
	}
	multiplyer := 1
	candidate := divisor * multiplyer
	for candidate < limit {
		multiples = append(multiples, candidate)
		multiplyer += 1
		candidate = divisor * multiplyer
	}
	return multiples
}
