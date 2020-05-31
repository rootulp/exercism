package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should return an abbreviated string based on s.
func Abbreviate(s string) string {
	parsed := replaceNonLettersWithWhitespace(s)

	fields := strings.Fields(parsed)
	result := ""
	for _, v := range fields {
		result += v[:1]
	}
	return strings.ToUpper(result)
}

func replaceNonLettersWithWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if r == '-' || r == ' ' {
			return ' '
		} else if unicode.IsLetter(r) {
			return r
		} else {
			return -1
		}
	}, s)
}
