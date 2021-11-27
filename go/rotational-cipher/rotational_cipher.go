package rotationalcipher

import (
	"log"
	"unicode"
)

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
	// 'a' is 97
	// 'z' is 122
	v := (int(r)-97+shiftKey)%26 + 97
	// log.Printf("rotateLower(%d, %d) = %d", r, shiftKey, v)
	return rune(v)
}

func rotateUpper(r rune, shiftKey int) rune {
	// 'A' is 65
	// 'Z' is 90
	v := (int(r)-65+shiftKey)%26 + 65
	log.Printf("rotateUpper(%d, %d) = %d", r, shiftKey, v)
	return rune(v)
}
