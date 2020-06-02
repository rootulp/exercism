package raindrops

import "strconv"

// Convert returns a string  generated based on the raindrops rules below.
// The rules of raindrops are that if a given number:
// - has 3 as a factor, add 'Pling' to the result.
// - has 5 as a factor, add 'Plang' to the result.
// - has 7 as a factor, add 'Plong' to the result.
// - does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
func Convert(n int) string {
	if !isDivisibleByPlingPlangPlong(n) {
		return strconv.Itoa(n)
	}
	var result string
	if isPling(n) {
		result += "Pling"
	}
	if isPlang(n) {
		result += "Plang"
	}
	if isPlong(n) {
		result += "Plong"
	}
	return result
}

func isDivisibleByPlingPlangPlong(n int) bool {
	return isPling(n) || isPlang(n) || isPlong(n)
}

func isPling(n int) bool {
	return n%3 == 0
}

func isPlang(n int) bool {
	return n%5 == 0
}

func isPlong(n int) bool {
	return n%7 == 0
}
