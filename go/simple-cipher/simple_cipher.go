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
	return NewShift(3).Encode(s)
}

func (c caesar) Decode(s string) (decoded string) {
	return NewShift(-3).Encode(s)
}

func NewShift(distance int) Cipher {
	if distance >= 26 || distance == 0 || distance <= -26 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) (encoded string) {
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			encoded += string(shiftLetter(r, c.distance))
		}
	}
	return encoded
}

func (c shift) Decode(input string) (decoded string) {
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			decoded += string(shiftLetter(r, -c.distance))
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

func shiftLetter(letter rune, distance int) (shifted rune) {
	// HACKHACK go % operator returns a negative number for -2 % 26 so add 26 then modulo 26
	// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
	d := rune(((int(letter-97)+distance)%26)+26)%26 + 97
	return rune(d)
}
