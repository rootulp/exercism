package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

func ExchangeCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[3:6], nil
}

func SubscriberNumber(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[6:], nil
}

func Number(phoneNumber string) (string, error) {
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

func clean(phoneNumber string) (cleaned string, err error) {
	for _, character := range phoneNumber {
		if unicode.IsDigit(character) {
			cleaned += string(character)
		}
	}
	if len(cleaned) == 11 && cleaned[0] == '1' {
		// trim leading 1
		cleaned = cleaned[1:]
	}
	if len(cleaned) != 10 {
		return "", fmt.Errorf("phone number %v must have 10 digits", cleaned)
	}
	return cleaned, err
}

func Format(phoneNumber string) (string, error) {
	areaCode, err := AreaCode(phoneNumber)
	if err != nil {
		return "", err
	}
	exchangeCode, err := ExchangeCode(phoneNumber)
	if err != nil {
		return "", err
	}
	subscriberNumber, err := SubscriberNumber(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%v) %v-%v", areaCode, exchangeCode, subscriberNumber), nil
}
