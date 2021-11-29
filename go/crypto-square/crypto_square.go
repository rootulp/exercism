package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(plainText string) (cipherText string) {
	stripped := stripFormatting(plainText)
	numCols, numRows := getRectangleDimensions(len(stripped))
	rectangle := getRectangle(stripped, numCols, numRows)
	encoded := getEncoded(rectangle, numCols, numRows)
	return strings.Join(splitEveryN(encoded, numRows), " ")
}

func stripFormatting(plainText string) (stripped string) {
	for _, r := range strings.ToLower(plainText) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			stripped += string(r)
		}
	}
	return stripped
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
	return rectangle
}

func initializeRectangle(numCols int, numRows int) (rectangle [][]rune) {
	rectangle = make([][]rune, numRows)
	for i := range rectangle {
		rectangle[i] = make([]rune, numCols)
	}
	return rectangle
}

func getEncoded(rectangle [][]rune, numCols int, numRows int) (encoded string) {
	for col := 0; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			encoded += string(rectangle[row][col])
		}
	}
	return encoded
}

func splitEveryN(message string, n int) (chunked []string) {
	for i := 0; i < len(message); i += n {
		upperBound := int(math.Min(float64(len(message)), float64(i+n)))
		chunked = append(chunked, message[i:upperBound])
	}
	return chunked
}
