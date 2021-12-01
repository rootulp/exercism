package cipher

import (
	"strings"
	"unicode"
)

type caesar struct{}
type shift struct {
	distance int
}
type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return caesar{}
}

func (c caesar) Encode(s string) (encoded string) {
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			encodedR := ((r-97)+3)%26 + 97
			encoded += string(encodedR)
		}
	}
	return encoded
}

func (c caesar) Decode(s string) (decoded string) {
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			// HACKHACK go % operator returns a negative number for -2 % 26 so add 26 then modulo 26
			// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
			decodedR := (((r-97)-3)%26+26)%26 + 97
			decoded += string(decodedR)
		}
	}
	return decoded
}

func NewShift(distance int) Cipher {
	return shift{distance: distance}
}

func (c shift) Encode(input string) (encoded string) {
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			encoded += string(shiftRune(r, c.distance))
		}
	}
	return encoded
}

func (c shift) Decode(input string) (decoded string) {
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			decoded += string(shiftRune(r, -c.distance))
		}
	}
	return decoded
}

func NewVigenere(key string) Cipher {
	return vigenere{
		key: key,
	}
}

func (v vigenere) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (v vigenere) Decode(input string) string {
	panic("Please implement the Decode function")
}

func shiftRune(r rune, shift int) (shiftedRune rune) {
	d := rune(((int(r-97)+shift)%26)+26)%26 + 97
	return rune(d)
}
