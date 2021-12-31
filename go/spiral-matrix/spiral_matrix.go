package spiralmatrix

type direction string

const east direction = "east"
const south direction = "south"
const west direction = "west"
const north direction = "north"

func SpiralMatrix(size int) (result [][]int) {
	if size == 0 {
		return [][]int{}
	}
	if size == 1 {
		return [][]int{{1}}
	}

	result = initMatrix(size)
	numToInsert := 1
	direction := east
	topRow := 0
	bottomRow := size - 1
	leftColumn := 0
	rightColumn := size - 1

	for topRow <= bottomRow && leftColumn <= rightColumn {
		switch direction {
		case east:
			for column := leftColumn; column <= rightColumn; column += 1 {
				result[topRow][column] = numToInsert
				numToInsert++
			}
			direction = advanceClockwise(direction)
			topRow += 1
		case south:
			for row := topRow; row <= bottomRow; row += 1 {
				result[row][rightColumn] = numToInsert
				numToInsert++
			}
			direction = advanceClockwise(direction)
			rightColumn -= 1
		case west:
			for column := rightColumn; column >= leftColumn; column -= 1 {
				result[bottomRow][column] = numToInsert
				numToInsert++
			}
			direction = advanceClockwise(direction)
			bottomRow -= 1
		case north:
			for row := bottomRow; row >= topRow; row -= 1 {
				result[row][leftColumn] = numToInsert
				numToInsert++
			}
			direction = advanceClockwise(direction)
			leftColumn += 1
		}
	}
	return result
}

func initMatrix(size int) (matrix [][]int) {
	matrix = [][]int{}
	for i := 0; i < size; i += 1 {
		row := make([]int, size)
		matrix = append(matrix, row)
	}
	return matrix
}

func advanceClockwise(dir direction) direction {
	switch dir {
	case east:
		return south
	case south:
		return west
	case west:
		return north
	case north:
		return east
	default:
		return "invalid direction"
	}
}
