package isogram

import "strings"

// IsIsogram returns whether the provided string is an isogram.
// In other words, whether the string does not contain any duplicate characters.
func IsIsogram(s string) bool {
	parsed := strings.ToLower(removeWhitespaceAndHyphens(s))
	seen := make(map[rune]bool)
	for _, c := range parsed {
		if (seen[c]) == true {
			return false
		}
		seen[c] = true
	}
	return true
}

func removeWhitespaceAndHyphens(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' || r == '-' {
			return -1
		}
		return r
	}, s)
}
