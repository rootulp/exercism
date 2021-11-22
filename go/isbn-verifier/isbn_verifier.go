package isbn

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	withoutDashes := strings.ReplaceAll(isbn, "-", "")
	if !isValidISBN(withoutDashes) {
		return false
	}
	return sum(withoutDashes)%11 == 0
}

func sum(isbn string) (sum int) {
	for index, digit := range isbn {
		value := valueForDigit(digit)
		multiplier := 10 - index
		sum += value * multiplier
	}
	return sum
}

func valueForDigit(digit rune) int {
	if digit == 'X' {
		return 10
	}
	value, err := strconv.Atoi(string(digit))
	if err != nil {
		log.Fatalf("failed to convert %v to int", digit)
	}
	return value
}

func isValidISBN(isbn string) bool {
	return isValidLength(isbn) && isValidCheckDigit(isbn) && isValidPrefix(isbn)
}

func isValidLength(isbn string) bool {
	return len(isbn) == 10
}

func isValidCheckDigit(isbn string) bool {
	if len(isbn) < 1 {
		return false
	}
	r := rune(isbn[len(isbn)-1])
	return unicode.IsDigit(r) || r == 'X'
}

// Prefix in this context is the entire ISBN excluding the check digit
func isValidPrefix(isbn string) bool {
	for _, c := range isbn[0 : len(isbn)-1] {
		if c != '-' && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
