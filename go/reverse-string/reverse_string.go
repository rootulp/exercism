package reverse

import (
	"strings"
)

// Reverse returns a string with characters in reversed order.
func Reverse(s string) string {
	chars := strings.Split(s, "")
	result := make([]string, len(s))

	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	return strings.Join(result, "")
}
