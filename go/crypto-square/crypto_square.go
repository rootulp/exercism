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
	encoded := getEncoded(rectangle)
	cipherText = strings.Join(splitEveryN(encoded, numRows), " ")
	return cipherText
}

func removeFormating(plainText string) (result string) {
	for _, r := range strings.ToLower(plainText) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
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
		return int(x + 1), int(x + 1)
	}
}

func getRectangle(message string, numCols int, numRows int) (rectangle [][]rune) {
	rectangle = initializeRectangle(numCols, numRows)
	index := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if index >= len(message) {
				rectangle[row][col] = ' '
			} else {
				rectangle[row][col] = []rune(message)[index]
				index += 1
			}
		}
	}
	fmt.Printf("rectangle %v\n", rectangle)
	return rectangle
}

func initializeRectangle(numCols int, numRows int) (rectangle [][]rune) {
	rectangle = make([][]rune, numRows)
	for i := range rectangle {
		rectangle[i] = make([]rune, numCols)
	}
	return rectangle
}

func getEncoded(rectangle [][]rune) (encoded string) {
	numRows := len(rectangle)
	numCols := len(rectangle[0])

	for col := 0; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			encoded += string(rectangle[row][col])
		}
	}
	fmt.Printf("encoded %v\n", encoded)
	return encoded
}

func splitEveryN(message string, n int) (chunked []string) {
	for i := 0; i < len(message); i += n {
		upperBound := int(math.Min(float64(len(message)), float64(i+n)))
		chunked = append(chunked, message[i:upperBound])
	}
	fmt.Printf("chunked %v\n", chunked)
	return chunked
}
