package transpose

import (
	"fmt"
	"strings"
)

func Transpose(input []string) (result []string) {
	if len(input) == 0 {
		return []string{}
	}
	transposed := make([][]string, len(input[0]))
	for _, v := range input {
		for col, c := range v {
			if transposed[col] == nil {
				transposed[col] = []string{}
			}
			fmt.Printf("transposed[%v]=append(%v, %v)\n", col, transposed[col], c)
			transposed[col] = append(transposed[col], string(c))
		}
	}
	fmt.Printf("transposed: %v\n", transposed)
	for _, row := range transposed {
		result = append(result, strings.Join(row, ""))
	}
	fmt.Printf("Transpose(%v)=%v\n", input, result)
	return result
}
