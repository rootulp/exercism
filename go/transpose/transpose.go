package transpose

import (
	"fmt"
	"strings"
)

func Transpose(input []string) (result []string) {
	if len(input) == 0 {
		return []string{}
	}
	transposed := make([][]string, maxLineLength(input))
	paddedInput := padInput(input)
	for _, v := range paddedInput {
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

func padInput(input []string) (padded []string) {
	for i, line := range input {
		lineLength := maxLineLength(input[i:])
		paddedLine := pad(line, lineLength)
		padded = append(padded, paddedLine)
	}
	fmt.Printf("pad(%v)=%v", input, padded)
	return padded
}

// pad right pads line with whitespace to a length of lineLength
func pad(line string, lineLength int) (padded string) {
	if len(line) >= lineLength {
		return line
	}
	return line + strings.Repeat(" ", lineLength-len(line))
}

func maxLineLength(input []string) (max int) {
	for _, line := range input {
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}
