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
	if unicode.IsLower(r) {
		return rotateLower(r, shiftKey)
	} else if unicode.IsUpper(r) {
		return rotateUpper(r, shiftKey)
	}
	return r
}

func rotateUpper(r rune, shiftKey int) rune {
	return rotateLetter(r, shiftKey, UPPERCASE_A)
}

func rotateLower(r rune, shiftKey int) rune {
	return rotateLetter(r, shiftKey, LOWERCASE_A)
}

func rotateLetter(r rune, shiftKey int, offset int) rune {
	return rune((int(r)-offset+shiftKey)%LETTERS_IN_ALPHABET + offset)
}
