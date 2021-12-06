package armstrong

import (
	"fmt"
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	return n == armstrongSum(n)
}

func armstrongSum(n int) (sum int) {
	digits := fmt.Sprint(n)
	for _, digit := range digits {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err)
		}
		sum += intPow(d, len(digits))
	}
	return sum
}

func intPow(n int, power int) int {
	return int(math.Pow(float64(n), float64(power)))
}
