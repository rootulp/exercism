package acronym

import "strings"

// Abbreviate should return an abbreviated string based on s.
func Abbreviate(s string) string {
	fields := strings.Fields(s)
	result := ""
	for _, v := range fields {
		result += v[:1]
	}
	return strings.ToUpper(result)
}
