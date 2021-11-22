package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
	if !isValidISBN(isbn) {
		return false
	}
	return false
}

func isValidISBN(isbn string) bool {
	if !isValidCheckDigit(rune(isbn[len(isbn)-1])) {
		return false
	}
	if !isValidPrefix(isbn) {
		return false
	}
	return true

}

func isValidCheckDigit(checkDigit rune) bool {
	return unicode.IsDigit(checkDigit) || checkDigit == 'X'
}

// Prefix in this context is the entire ISBN excluding the check digit
func isValidPrefix(isbn string) bool {
	for _, c := range isbn[0 : len(isbn)-1] {
		if c != '-' || unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
