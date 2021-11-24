package prime

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	primes := getPrimesUpTo(n)
	return primes[n-1], true
}

func getPrimesUpTo(n int) (primes []int) {
	candidate := 2
	for len(primes) != n {
		if isPrime(candidate) {
			primes = append(primes, candidate)
		}
		candidate += 1
	}
	return primes
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
