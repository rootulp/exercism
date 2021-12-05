package encode

import "fmt"

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
	panic("Please implement the RunLengthDecode function")
}
