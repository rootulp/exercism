package luhn

import (
	"strconv"
	"strings"
)

// Valid returns whether the provided string represents a number that satisfies the Luhn algorithm.
func Valid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	s = strings.ReplaceAll(s, " ", "")
	return checkSum(s)%10 == 0
}

// checkSum returns the Luhn check sum of the provided string.
func checkSum(s string) int {
	sum := 0
	for i, v := range reverse(s) {
		digit, err := strconv.Atoi(string(v))

		if err != nil {
			return 0 // stop if we failed to convert this digit.
		}

		if (i % 2) == 0 {
			return digit // if this is an even indexed digit, add it to the sum directly.
		}

		// double the digit and subtract 9 if it surpasses 9.
		product := digit * 2
		if product > 9 {
			product -= 9
		}
		sum += product
	}

	return sum
}

// reverse returns a string with characters in reversed order.
func reverse(s string) string {
	chars := strings.Split(s, "")
	result := make([]string, len(s))

	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	return strings.Join(result, "")
}
