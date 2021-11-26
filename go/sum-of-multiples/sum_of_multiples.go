package summultiples

import "fmt"

func SumMultiples(limit int, divisors ...int) (sum int) {
	multiples := getMultiples(limit, divisors)
	for multiple := range multiples {
		sum += multiple
	}
	return sum
}

func getMultiples(limit int, divisors []int) (multiples map[int]bool) {
	multiples = map[int]bool{}
	for _, divisor := range divisors {
		multiplesForDivisor := getMultiplesForDivisor(limit, divisor)
		for _, multiple := range multiplesForDivisor {
			multiples[multiple] = true
		}
	}
	return multiples
}

func getMultiplesForDivisor(limit int, divisor int) (multiples []int) {
	if divisor == 0 {
		return multiples
	}
	multiplyer := 1
	candidate := divisor * multiplyer
	for candidate < limit {
		fmt.Printf("candidate %d, multiplyer %d, limit %d\n", candidate, multiplyer, limit)
		multiples = append(multiples, candidate)
		multiplyer += 1
		candidate = divisor * multiplyer
	}
	fmt.Printf("returning multiples %v\n", multiples)
	return multiples
}
