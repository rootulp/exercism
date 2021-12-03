package allyourbase

import (
	"fmt"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (outputDigits []int, e error) {
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

func reverse(input []int) (reversed []int) {
	for i := len(input) - 1; i >= 0; i-- {
		reversed = append(reversed, input[i])
	}
	return reversed
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
