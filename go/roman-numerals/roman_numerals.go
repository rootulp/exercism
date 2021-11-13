package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 {
		return "", errors.New("input must be greater than 0")
	} else if input > 3000 {
		return "", errors.New("input must be less than or equal to 3000")
	}
	output := convertDigitToRomanNumeral(input)
	return output, nil
}

func convertDigitToRomanNumeral(digit int) string {
	digitToRomanNumeral := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}
	return digitToRomanNumeral[digit]
}
