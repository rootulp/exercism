package transpose

import (
	"strings"
)

func Transpose(input []string) (result []string) {
	result = []string{}
	transposed := make([][]string, maxLineLength(input))
	for _, row := range padInput(input) {
		for col, v := range row {
			if transposed[col] == nil {
				transposed[col] = []string{}
			}
			transposed[col] = append(transposed[col], string(v))
		}
	}
	for _, row := range transposed {
		result = append(result, strings.Join(row, ""))
	}
	return result
}

// padInput pads input so that each line has a length equal to the max line
// length of all lines that follow
func padInput(input []string) (padded []string) {
	for i, line := range input {
		lineLength := maxLineLength(input[i:])
		paddedLine := pad(line, lineLength)
		padded = append(padded, paddedLine)
	}
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
