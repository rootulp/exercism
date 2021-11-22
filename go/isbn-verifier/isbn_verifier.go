package isbn

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbnWithoutDashes := strings.ReplaceAll(isbn, "-", "")
	if !isValidISBN(isbnWithoutDashes) {
		return false
	}
	return isbnSum(isbnWithoutDashes)%11 == 0
}

func isbnSum(isbn string) (sum int) {
	for i, r := range isbn {
		value := valueForCheckDigit(string(r))
		multiplier := 10 - i
		sum += value * multiplier
	}
	log.Printf("isbnSum(%s): %d\n", isbn, sum)
	return sum
}

func valueForCheckDigit(r string) int {
	if r == "X" || r == "x" {
		return 10
	}
	value, err := strconv.Atoi(r)
	if err != nil {
		log.Fatalf("failed to convert %v to int", r)
	}
	return value
}

func isValidISBN(isbn string) bool {
	return isValidLength(isbn) && isValidCheckDigit(isbn) && isValidPrefix(isbn)
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

func isValidLength(isbn string) bool {
	return len(isbn) == 10
}
