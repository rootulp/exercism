package cipher

import (
	"strings"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	distance int
}
type vigenere struct{}
type Caesar struct{}

func NewCaesar() Cipher {
	return Caesar{}
}

func (c Caesar) Encode(s string) (encoded string) {
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			encodedR := ((r-97)+3)%26 + 97
			encoded += string(encodedR)
		}
	}
	return encoded
}

func (c Caesar) Decode(s string) (decoded string) {
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
			encodedR := rune(int(r-97)+c.distance)%26 + 97
			encoded += string(encodedR)
		}
	}
	return encoded
}

func (c shift) Decode(input string) string {
	panic("Please implement the Decode function")
}

func NewVigenere(key string) Cipher {
	panic("Please implement the NewVigenere function")
}

func (v vigenere) Encode(input string) string {
	panic("Please implement the Encode function")
}

func (v vigenere) Decode(input string) string {
	panic("Please implement the Decode function")
}
