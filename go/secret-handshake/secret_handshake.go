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
		if v == '1' {
			handshake = append(handshake, operations[i])
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
