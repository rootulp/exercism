package rotationalcipher

import (
	"unicode"
)

const LETTERS_IN_ALPHABET = 26
const UPPERCASE_A = 65
const LOWERCASE_A = 97

func RotationalCipher(plain string, shiftKey int) (cipherText string) {
	for _, r := range plain {
		cipherText += string(rotate(r, shiftKey))
	}
	return cipherText
}

func rotate(r rune, shiftKey int) (result rune) {
	if unicode.IsPunct(r) {
		return r
	}
	if unicode.IsLower(r) {
		return rotateLower(r, shiftKey)
	} else if unicode.IsUpper(r) {
		return rotateUpper(r, shiftKey)
	}
	return r
}

func rotateLower(r rune, shiftKey int) rune {
	v := (int(r)-LOWERCASE_A+shiftKey)%LETTERS_IN_ALPHABET + LOWERCASE_A
	return rune(v)
}

func rotateUpper(r rune, shiftKey int) rune {
	v := (int(r)-UPPERCASE_A+shiftKey)%LETTERS_IN_ALPHABET + UPPERCASE_A
	return rune(v)
}
