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
		ciphertext += string(encode(r))
	}
	return strings.Join(splitEveryN(ciphertext, 5), " ")
}

func removeNonAlphanumeric(s string) (result string) {
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			result += string(r)
		}
	}
	return result
}

func splitEveryN(message string, n int) (groups []string) {
	for i := 0; i < len(message); i += n {
		upperBound := math.Min(float64(len(message)), float64(i+n))
		group := message[i:int(upperBound)]
		groups = append(groups, group)
	}
	return groups
}

func encode(r rune) rune {
	if unicode.IsLetter(r) {
		return atbash[r]
	}
	return r
}
