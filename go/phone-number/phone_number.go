package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (number string, e error) {
	for _, character := range phoneNumber {
		if unicode.IsDigit(character) {
			number += string(character)
		}
	}
	if len(number) == 11 && number[0] == '1' {
		// trim leading 1
		number = number[1:]
	}
	if len(number) != 10 {
		return "", fmt.Errorf("phone number %v must have 10 digits", number)
	}
	_, err := AreaCode(number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func AreaCode(phoneNumber string) (areaCode string, e error) {
	areaCode = phoneNumber[0:3]
	if strings.HasPrefix(areaCode, "0") || strings.HasPrefix(areaCode, "1") {
		return "", fmt.Errorf("area code %v can not start with 0", areaCode)
	}
	return areaCode, nil
}

func Format(phoneNumber string) (string, error) {
	panic("Please implement the Format function")
}
