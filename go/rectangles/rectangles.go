package rectangles

const CORNER rune = '+'
const VERTICAL_EDGE rune = '|'
const HORIZONTAL_EDGE rune = '-'

func Count(diagram []string) (rectangles int) {
	for rowIndex, row := range diagram {
		for columnIndex, character := range row {
			if isCorner(character) {
				validColumns := findValidColumns(diagram, columnIndex, rowIndex)
				validRows := findValidRows(diagram, columnIndex, rowIndex)
				rectangles += countRectanglesForTopLeft(diagram, validColumns, validRows, columnIndex, rowIndex)
			}
		}
	}
	return rectangles
}

// Finds columns and rows that could contain the other corners
// for a rectangle with a particular point as the top left corner
func findValidColumns(diagram []string, columnIndex int, rowIndex int) (validColumns []int) {
	validColumns = []int{}
	row := diagram[rowIndex]

	for i := columnIndex + 1; i < len(row); i++ {
		if isCorner(rune(row[i])) {
			validColumns = append(validColumns, i)
		} else if !isHorizontalEdge(rune(row[i])) {
			return validColumns
		}
	}
	return validColumns
}

func findValidRows(diagram []string, columnIndex int, rowIndex int) (validRows []int) {
	validRows = []int{}

	for j := rowIndex + 1; j < len(diagram); j++ {
		if isCorner(rune(diagram[j][columnIndex])) {
			validRows = append(validRows, j)
		} else if !isVerticalEdge(rune(diagram[j][columnIndex])) {
			return validRows
		}
	}
	return validRows
}

// Counts the number of rectangles for a particular point being
// the top left corner
func countRectanglesForTopLeft(diagram []string, validColumns, validRows []int, columnIndex, rowIndex int) (rectangles int) {
	for _, k := range validColumns {
		for _, l := range validRows {
			brokenRectangle := false
			for m := columnIndex + 1; m < k; m++ {
				if diagram[l][m] != '+' && diagram[l][m] != '-' {
					brokenRectangle = true
					break
				}
			}
			for n := rowIndex + 1; n < l; n++ {
				if diagram[n][k] != '+' && diagram[n][k] != '|' {
					brokenRectangle = true
					break
				}
			}
			if diagram[l][k] == '+' && !brokenRectangle {
				rectangles++
			}
		}
	}
	return rectangles
}

func isCorner(r rune) bool {
	return r == CORNER
}

func isHorizontalEdge(r rune) bool {
	return r == HORIZONTAL_EDGE
}

func isVerticalEdge(r rune) bool {
	return r == VERTICAL_EDGE
}
