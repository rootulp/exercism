package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 {
		return "", errors.New("input must be greater than 0")
	} else if input >= 3000 {
		return "", errors.New("input must be less than or equal to 3000")
	}
	return "0", nil
}
