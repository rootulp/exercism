package sieve

func Sieve(limit int) (primes []int) {
	if limit < 2 {
		return []int{}
	}
	unprocessed := createRange(2, limit)
	for len(unprocessed) != 0 {
		prime := unprocessed[0]
		unprocessed = unprocessed[1:]

		unprocessed = removeMultiples(unprocessed, prime)
		primes = append(primes, prime)
	}
	return primes
}

func removeMultiples(slice []int, prime int) (filtered []int) {
	if len(slice) == 0 {
		return slice
	}
	multiplyer := 2
	multiple := multiplyer * prime
	for len(slice) != 0 && multiple <= slice[len(slice)-1] {
		slice = remove(slice, multiple)
		multiplyer += 1
		multiple = multiplyer * prime
	}
	return slice
}

func remove(slice []int, value int) (result []int) {
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func createRange(lower int, upper int) (result []int) {
	for i := lower; i <= upper; i++ {
		result = append(result, i)
	}
	return result
}
