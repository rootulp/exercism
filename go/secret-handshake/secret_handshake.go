package secret

import (
	"fmt"
	"strconv"
)

var operations = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) (handshake []string) {
	binary := strconv.FormatInt(int64(code), 2)
	fmt.Printf("code %v in binary %v\n", code, binary)
	for i, v := range reverse(binary) {
		if v == '1' && i < len(operations) {
			handshake = append(handshake, operations[i])
		}
		if v == '1' && i == len(operations) {
			handshake = reverseSlice(handshake)
		}
	}
	return handshake
}

func reverse(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return result
}

func reverseSlice(input []string) (result []string) {
	for _, v := range input {
		result = append([]string{v}, result...)
	}
	return result
}
