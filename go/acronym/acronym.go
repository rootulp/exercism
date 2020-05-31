package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should return an abbreviated string based on s.
func Abbreviate(s string) string {
	parsed := parse(s)
	acronym := generateAcronym(parsed)
	return strings.ToUpper(acronym)
}

// Return a string that is generated from the first letter of each word in s.
func generateAcronym(s string) string {
	result := ""
	for _, v := range strings.Fields(s) {
		result += v[:1]
	}
	return result
}

func parse(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '-' || r == ' ' {
			return ' ' // Convert hyphens to whitespace
		} else if unicode.IsLetter(r) {
			return r
		} else {
			return -1 // Remove non whitespace or letter characters
		}
	}, s)
}
