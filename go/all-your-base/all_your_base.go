package allyourbase

import (
	"errors"
	"fmt"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (outputDigits []int, e error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}
	if invalidDigit(inputBase, inputDigits) {
		return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
	}
	base10 := getBase10Input(inputBase, inputDigits)
	if base10 == 0 {
		return []int{0}, nil
	}
	for base10 > 0 {
		digit := base10 % outputBase
		outputDigits = append([]int{digit}, outputDigits...)
		base10 = base10 / outputBase
	}
	return outputDigits, nil
}

func getBase10Input(inputBase int, inputDigits []int) (base10Input int) {
	for i, digit := range reverse(inputDigits) {
		base10Input += powInt(inputBase, i) * digit
	}
	fmt.Printf("getBase10Input(%d, %v)=%d\n", inputBase, inputDigits, base10Input)
	return base10Input
}

func invalidDigit(inputBase int, input []int) bool {
	for _, digit := range input {
		if digit < 0 {
			return true
		}
		if digit >= inputBase {
			return true
		}
	}
	return false
}

func reverse(input []int) (reversed []int) {
	for i := len(input) - 1; i >= 0; i-- {
		reversed = append(reversed, input[i])
	}
	return reversed
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
