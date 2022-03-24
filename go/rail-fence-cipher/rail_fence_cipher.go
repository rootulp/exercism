package railfence

import "fmt"

const Down = "down"
const Up = "up"

func Encode(message string, numRails int) (encoded string) {
	rails := initRails(numRails)
	railIndex := 0
	direction := Down

	for _, c := range message {
		rails[railIndex] = rails[railIndex] + string(c)
		direction, railIndex = advance(direction, railIndex, numRails)
	}

	fmt.Printf("rails: %v\n", rails)
	for _, rail := range rails {
		encoded += rail
	}
	return encoded
}

func Decode(message string, rails int) (decoded string) {
	return ""
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
