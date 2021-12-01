package cipher

import (
	"fmt"
	"strings"
	"unicode"
)

type caesar struct {
	distance int
}
type shift struct {
	distance int
}
type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return caesar{
		distance: 3,
	}
}

func (c caesar) Encode(s string) (encoded string) {
	return NewShift(c.distance).Encode(s)
}

func (c caesar) Decode(s string) (decoded string) {
	return NewShift(-c.distance).Encode(s)
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
	if !isValid(key) {
		return nil
	}
	return vigenere{
		key: key,
	}
}

func (v vigenere) Encode(input string) (encoded string) {
	for i, r := range stripFormatting(input) {
		distance := getDistance(v.key, i)
		if unicode.IsLetter(r) {
			fmt.Printf("r: %v, distance: %v, shifted: %v\n", string(r), distance, string(shiftLetter(r, distance)))
			encoded += string(shiftLetter(r, distance))
		}
	}
	return encoded
}

func (v vigenere) Decode(input string) (decoded string) {
	for i, r := range stripFormatting(input) {
		distance := getDistance(v.key, i)
		if unicode.IsLetter(r) {
			decoded += string(shiftLetter(r, -distance))
		}
	}
	return decoded
}

func shiftLetter(letter rune, distance int) (shifted rune) {
	// HACKHACK go % operator returns a negative number for -2 % 26 so add 26 then modulo 26
	// https://stackoverflow.com/questions/43018206/modulo-of-negative-integers-in-go
	d := rune(((int(letter-97)+distance)%26)+26)%26 + 97
	return rune(d)
}

func getDistance(key string, index int) (distance int) {
	position := index % len(key)
	distance = int([]rune(key)[position] - 97)
	fmt.Printf("key %v, position %v, distance %v\n", key, position, distance)
	return distance
}

func stripFormatting(input string) (stripped string) {
	for _, c := range strings.ToLower(input) {
		if unicode.IsLetter(c) {
			stripped += string(c)
		}
	}
	return stripped
}

func isValid(key string) bool {
	if key == "" {
		return false
	}
	if key == "a" || key == "aa" {
		return false
	}
	if stripFormatting(key) != key {
		return false
	}
	return true
}
