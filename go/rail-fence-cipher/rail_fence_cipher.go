package railfence

import (
	"fmt"
	"unicode/utf8"
)

const Down = "down"
const Up = "up"

func Encode(message string, numRails int) (encoded string) {
	rails := encodeRails(message, numRails)
	for _, rail := range rails {
		encoded += rail
	}
	return encoded
}

func encodeRails(message string, numRails int) []string {
	rails := initRails(numRails)
	railIndex := 0
	direction := Down

	for _, c := range message {
		rails[railIndex] = rails[railIndex] + string(c)
		direction, railIndex = advance(direction, railIndex, numRails)
	}

	fmt.Printf("encodeRails: %v\n", rails)
	return rails
}

func Decode(message string, numRails int) (decoded string) {
	rails := splitMessageIntoRails(message, numRails)
	railIndex := 0
	direction := Down

	for rails[railIndex] != "" {
		r, trimmed := trimFirstRune(rails[railIndex])
		decoded += string(r)
		rails[railIndex] = trimmed
		direction, railIndex = advance(direction, railIndex, numRails)
	}

	return decoded
}

func trimFirstRune(s string) (rune, string) {
	r, i := utf8.DecodeRuneInString(s)
	return r, s[i:]
}

func splitMessageIntoRails(message string, numRails int) []string {
	copy := message
	rails := initRails(numRails)

	for i := range rails {
		length := lengthOfRail(message, numRails, i)
		rails[i] = copy[:length]
		copy = copy[length:]
		// fmt.Printf("index %v, rail %v, length %v, copy %v\n", i, rails[i], length, copy)
	}

	return rails
}

func lengthOfRail(message string, numRails int, index int) int {
	rails := encodeRails(message, numRails)
	return len(rails[index])
}

func initRails(numRails int) (rails []string) {
	for i := 0; i < numRails; i++ {
		rails = append(rails, "")
	}
	return rails
}

func advance(direction string, railIndex int, numRails int) (nextDirection string, nextIndex int) {
	switch direction {
	case "down":
		if railIndex == numRails-1 { // reached bottom, turn around
			return "up", railIndex - 1
		}
		return "down", railIndex + 1
	case "up":
		if railIndex == 0 { // reached bottom, turn around
			return "down", railIndex + 1
		}
		return "up", railIndex - 1
	default:
		panic(fmt.Sprintf("invalid direction %v", direction))
	}
}
