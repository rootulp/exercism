package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns whether the provided string is an isogram.
// In other words, whether the string does not contain any duplicate characters.
func IsIsogram(s string) bool {
	parsed := strings.ToLower(preserveOnlyLetters(s))
	seen := make(map[rune]bool)
	for _, c := range parsed {
		if (seen[c]) == true {
			return false
		}
		seen[c] = true
	}
	return true
}

func preserveOnlyLetters(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, s)
}
