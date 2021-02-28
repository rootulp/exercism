package luhn

import (
	"errors"
	"strconv"
	"strings"
)

// Valid returns whether the provided string represents a number that satisfies the Luhn algorithm.
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) <= 1 {
		return false
	}
	checkDigit, err := checkSum(s)

	if err != nil {
		return false
	}

	return checkDigit%10 == 0
}

// checkSum returns the Luhn check sum of the provided string.
func checkSum(s string) (int, error) {
	sum := 0
	for i, v := range reverse(s) {
		digit, err := strconv.Atoi(string(v))

		if err != nil {
			// stop if we failed to convert this digit
			return 0, errors.New("Invalid digit")
		}

		if (i % 2) == 0 {
			// if this is an even indexed digit, add it to the sum directly
			sum += digit
		} else {
			// double the digit and subtract 9 if it surpasses 9
			product := digit * 2
			if product > 9 {
				product -= 9
			}
			sum += product
		}
	}
	return sum, nil
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
