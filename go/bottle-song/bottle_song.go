package bottlesong

import (
	"fmt"
	"strings"
)

var numberToCardinal = map[int]string{
	10: "ten",
	9:  "nine",
	3:  "three",
	2:  "two",
}

func Recite(startBottles, takeDown int) []string {
	startCardinal := cardinal(startBottles)
	endCardinal := cardinal(startBottles - 1)

	return []string{
		fmt.Sprintf("%v green bottles hanging on the wall,", capitalize(startCardinal)),
		fmt.Sprintf("%v green bottles hanging on the wall,", capitalize(startCardinal)),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %v green bottles hanging on the wall.", endCardinal),
	}
}

func cardinal(number int) string {
	return numberToCardinal[number]
}

// capitalize capitalizes the first letter of a given string
func capitalize(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}
