package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

func AreaCode(phoneNumber string) (areaCode string, e error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

func Number(phoneNumber string) (number string, e error) {
	number, err := clean(phoneNumber)
	if err != nil {
		return "", err
	}
	areaCode := number[0:3]
	if strings.HasPrefix(areaCode, "0") || strings.HasPrefix(areaCode, "1") {
		return "", fmt.Errorf("area code %v can not start with 0 or 1", areaCode)
	}
	exchangeCode := number[3:6]
	if strings.HasPrefix(exchangeCode, "0") || strings.HasPrefix(exchangeCode, "1") {
		return "", fmt.Errorf("exchange code %v can not start with 0 or 1", exchangeCode)
	}
	return number, nil
}

func clean(phoneNumber string) (number string, e error) {
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
	return number, e
}

func Format(phoneNumber string) (string, error) {
	panic("Please implement the Format function")
}
