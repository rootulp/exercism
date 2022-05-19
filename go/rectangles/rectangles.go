package rectangles

func Count(diagram []string) int {
	counter := 0
	for rowIndex, row := range diagram {
		for columnIndex, char := range row {
			if char == '+' {
				validColumns, validRows := FindValidColumnsAndRows(diagram, columnIndex, rowIndex)
				counter += CountRectanglesForTopLeft(diagram, validColumns, validRows, columnIndex, rowIndex)
			}
		}
	}
	return counter
}

// Finds columns and rows that could contain the other corners
// for a rectangle with a particular point as the top left corner
func FindValidColumnsAndRows(diagram []string, columnIndex, rowIndex int) ([]int, []int) {
	row := diagram[rowIndex]
	validColumns := []int{}
	validRows := []int{}
	for i := columnIndex + 1; i < len(diagram[rowIndex]); i++ {
		if row[i] == '+' {
			validColumns = append(validColumns, i)
		} else if row[i] != '-' {
			break
		}
	}
	for j := rowIndex + 1; j < len(diagram); j++ {
		if diagram[j][columnIndex] == '+' {
			validRows = append(validRows, j)
		} else if diagram[j][columnIndex] != '|' {
			break
		}
	}
	return validColumns, validRows
}

// Counts the number of rectangles for a particular point being
// the top left corner
func CountRectanglesForTopLeft(diagram []string, validColumns, validRows []int, columnIndex, rowIndex int) int {
	counter := 0
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
				counter++
			}
		}
	}
	return counter
}
