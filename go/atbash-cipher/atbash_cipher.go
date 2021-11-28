package atbash

import (
	"math"
	"strings"
	"unicode"
)

// Plain:  abcdefghijklmnopqrstuvwxyz
// Cipher: zyxwvutsrqponmlkjihgfedcba
var atbash = map[rune]rune{
	'a': 'z',
	'b': 'y',
	'c': 'x',
	'd': 'w',
	'e': 'v',
	'f': 'u',
	'g': 't',
	'h': 's',
	'i': 'r',
	'j': 'q',
	'k': 'p',
	'l': 'o',
	'm': 'n',
	'n': 'm',
	'o': 'l',
	'p': 'k',
	'q': 'j',
	'r': 'i',
	's': 'h',
	't': 'g',
	'u': 'f',
	'v': 'e',
	'w': 'd',
	'x': 'c',
	'y': 'b',
	'z': 'a',
}

func Atbash(message string) (ciphertext string) {
	for _, r := range removeNonAlphanumeric(strings.ToLower(message)) {
		if unicode.IsDigit(r) {
			ciphertext += string(r)
		} else {
			ciphertext += string(atbash[r])
		}
	}
	return splitEveryN(ciphertext, 5)
}

func removeNonAlphanumeric(s string) (result string) {
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			result += string(r)
		}
	}
	return result
}

func splitEveryN(message string, n int) (result string) {
	arr := []string{}
	for i := 0; i < len(message); i += n {
		upperBound := math.Min(float64(len(message)), float64(i+n))
		arr = append(arr, message[i:int(upperBound)])
	}
	return strings.Join(arr, " ")
}
