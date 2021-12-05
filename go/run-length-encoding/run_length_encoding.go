package encode

import (
	"fmt"
	"strconv"
	"unicode"
)

func RunLengthEncode(input string) (encoded string) {
	if input == "" {
		return input
	}
	var prevChar rune
	prevOccurences := 0
	for _, char := range input {
		if prevChar == 0 { // an uninitizlied rune is 0
			prevChar = char
			prevOccurences += 1
		} else if char != prevChar {
			encoded += encodeSegment(prevChar, prevOccurences)
			prevChar = char
			prevOccurences = 1
		} else if char == prevChar {
			prevOccurences += 1
		}
	}
	encoded += encodeSegment(prevChar, prevOccurences)
	return encoded
}

func encodeSegment(char rune, occurences int) string {
	if occurences == 1 {
		return string(char)
	}
	return fmt.Sprintf("%d%v", occurences, string(char))
}

func RunLengthDecode(input string) (decoded string) {
	if input == "" {
		return input
	}

	runes := []rune(input)
	for i := 0; i < len(runes); i += 1 {
		if unicode.IsLetter(runes[i]) {
			decoded += string(runes[i])
		} else if unicode.IsDigit(runes[i]) {
			occurences, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				panic(err)
			}
			decoded += decodeSegment(runes[i+1], occurences)
			i += 1
		}
	}
	return decoded
}

func decodeSegment(character rune, occurences int) (segment string) {
	for i := 0; i < occurences; i++ {
		segment += string(character)
	}
	return segment
}
