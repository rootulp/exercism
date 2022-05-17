package diamond

import (
	"fmt"
	"strings"
)

func Gen(char byte) (result string, err error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("char %v is not a valid capital letter", char)
	}
	if char == 'A' {
		return "A\n", nil
	}

	dimension := getDimension(rune(char))
	// fmt.Printf("dimension %v\n", dimension)

	rows := []string{}
	rows = append(rows, header(dimension)...)
	rows = append(rows, middle(dimension, rune(char)))
	rows = append(rows, footer(dimension)...)
	return strings.Join(rows, "\n") + "\n", nil
}

func header(dimension int) (rows []string) {
	headerLength := headerLength(dimension)
	for rowNumber := 0; rowNumber < headerLength; rowNumber++ {
		row := row(dimension, rowNumber)
		rows = append(rows, row)
	}
	return rows
}

func row(dimension int, rowNumber int) (result string) {
	leadingSpaces := leadingSpaces(dimension, rowNumber)
	character := characterForRow(rowNumber)
	middleSpaces := middleSpaces(dimension, rowNumber)

	if rowNumber == 0 {
		fmt.Printf("rowNumber %v and leading spaces %v\n", rowNumber, leadingSpaces)
		return fmt.Sprintf("%s%s%s", strings.Repeat(" ", leadingSpaces), string(character), strings.Repeat(" ", leadingSpaces))

	}

	result += strings.Repeat(" ", leadingSpaces)
	result += string(character)
	result += strings.Repeat(" ", middleSpaces)
	result += string(character)
	result += strings.Repeat(" ", leadingSpaces)
	return result
}

func headerLength(dimension int) int {
	return (dimension - 1) / 2
}

func leadingSpaces(dimension int, rowNumber int) (numSpaces int) {
	return headerLength(dimension) - rowNumber
}

func characterForRow(rowNumber int) rune {
	return rune('A' + rowNumber)
}

func middleSpaces(dimension int, rowNumber int) (numSpaces int) {
	if rowNumber == 0 {
		return 0
	}
	return rowNumber*2 - 1
}

func middle(dimension int, r rune) (row string) {
	row += string(r)
	row += strings.Repeat(" ", middleSpaces(dimension, headerLength(dimension)))
	row += string(r)
	return row
}

func footer(dimension int) (rows []string) {
	return reverse(header(dimension))
}

func reverse(rows []string) (reversed []string) {
	for _, row := range rows {
		reversed = append([]string{row}, reversed...)
	}
	return reversed
}

func getDimension(char rune) (dimension int) {
	dimension = 1
	for char != 'A' {
		dimension += 2
		char -= 1
	}
	return dimension
}
