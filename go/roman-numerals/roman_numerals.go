package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (output string, err error) {
	if input <= 0 {
		return "", errors.New("input must be greater than 0")
	} else if input > 3000 {
		return "", errors.New("input must be less than or equal to 3000")
	}

	numerator := input
	denominator := 1000
	for numerator != 0 {
		quotient, remainder := divmod(numerator, denominator)

		output += convertToNumeral(quotient * denominator)

		numerator = remainder
		denominator = denominator / 10
	}
	return output, nil
}

func convertToNumeral(number int) string {
	numberToNumeral := map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		20:   "XX",
		30:   "XXX",
		40:   "XL",
		50:   "L",
		60:   "LX",
		70:   "LXX",
		80:   "LXXX",
		90:   "IC",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	return numberToNumeral[number]
}

func divmod(numerator int, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return quotient, remainder
}
