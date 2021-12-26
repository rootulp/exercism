package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be positive")
	}
	return collatz(n, 0), nil
}

func collatz(n int, steps int) int {
	if n == 1 {
		return steps
	}
	steps += 1
	if n%2 == 0 {
		return collatz(n/2, steps)
	} else {
		return collatz(3*n+1, steps)
	}
}
