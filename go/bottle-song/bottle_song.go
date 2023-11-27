package bottlesong

import (
	"fmt"
	"strings"
)

var numberToCardinal = map[int]string{
	10: "ten",
	9:  "nine",
	8:  "eight",
	7:  "seven",
	6:  "six",
	5:  "five",
	4:  "four",
	3:  "three",
	2:  "two",
	1:  "one",
	0:  "no",
}

func Recite(startBottles, takeDown int) (song []string) {
	for i := startBottles; i > startBottles-takeDown; i-- {
		song = append(song, verse(i)...)
		song = append(song, "")
	}
	return song[:len(song)-1] // trim the last ""
}

func verse(bottles int) []string {
	startCardinal := cardinal(bottles)
	endCardinal := cardinal(bottles - 1)

	return []string{
		fmt.Sprintf("%v green %v hanging on the wall,", capitalize(startCardinal), maybePluralize(bottles, "bottle", "bottles")),
		fmt.Sprintf("%v green %v hanging on the wall,", capitalize(startCardinal), maybePluralize(bottles, "bottle", "bottles")),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf("There'll be %v green %v hanging on the wall.", endCardinal, maybePluralize(bottles-1, "bottle", "bottles")),
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

// maybePluralize returns the plural version of word if number is > 1. Returns
// the singular version of word if number == 1.
func maybePluralize(number int, singularWord string, pluralWord string) string {
	if number == 1 {
		return singularWord
	}
	return pluralWord
}
