package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func Encode(plainText string) (cipherText string) {
	formatted := removeFormating(plainText)
	numCols, numRows := getRectangleDimensions(len(formatted))
	rectangle := getRectangle(formatted, numCols, numRows)
	fmt.Printf("rectangle %v\n", rectangle)
	return formatted
}

func removeFormating(plainText string) (result string) {
	for _, r := range strings.ToLower(plainText) {
		if unicode.IsLetter(r) {
			result += string(r)
		}
	}
	return result
}

func getRectangleDimensions(messageLength int) (numCols int, numRows int) {
	x := math.Trunc(math.Sqrt(float64(messageLength)))
	if x*x >= float64(messageLength) {
		return int(x), int(x)
	} else if (x+1)*x >= float64(messageLength) {
		return int(x + 1), int(x)
	} else {
		return int(x + 1), int(x)
	}
}

func getRectangle(message string, numCols int, numRows int) (rectangle [][]rune) {
	rectangle = initializeRectangle(numCols, numRows)
	index := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			rectangle[row][col] = []rune(message)[index]
			index += 1
		}
	}
	return rectangle
}

func initializeRectangle(numCols int, numRows int) (rectangle [][]rune) {
	rectangle = make([][]rune, numRows)
	for i := range rectangle {
		rectangle[i] = make([]rune, numCols)
	}
	return rectangle
}
