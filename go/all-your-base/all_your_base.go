package allyourbase

import (
	"fmt"
	"math"
	"strconv"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (outputDigits []int, e error) {
	base10 := getBase10Input(inputBase, inputDigits)
	for i, d := range reverse(getDigits(base10)) {
		digit := powInt(outputBase, i) * d
		outputDigits = append(outputDigits, digit)
	}
	return outputDigits, nil
}

func getDigits(input int) (digits []int) {
	for _, d := range strconv.Itoa(input) {
		digits = append([]int{int(d - '0')}, digits...)
	}
	fmt.Printf("digits(%d)=%v\n", input, digits)
	return digits
}

func getBase10Input(inputBase int, inputDigits []int) (base10Input int) {
	for i, digit := range reverse(inputDigits) {
		base10Input += powInt(inputBase, i) * digit
	}
	fmt.Printf("getBase10Input(%d, %v)=%d\n", inputBase, inputDigits, base10Input)
	return base10Input
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
