package wordsearch

import (
	"fmt"
)

const NORTH = "NORTH"
const EAST = "EAST"
const SOUTH = "SOUTH"
const WEST = "WEST"
const NORTH_EAST = "NORTH_EAST"
const NORTH_WEST = "NORTH_WEST"
const SOUTH_EAST = "SOUTH_EAST"
const SOUTH_WEST = "SOUTH_WEST"

var DIRECTIONS = []string{NORTH, EAST, SOUTH, WEST, NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST}
var NOT_FOUND_POINT = [2]int{-1, -1}
var NOT_FOUND_LINE = [2][2]int{NOT_FOUND_POINT, NOT_FOUND_POINT}

func Solve(words []string, puzzle []string) (result map[string][2][2]int, err error) {
	grid := NewGrid(puzzle)
	result = make(map[string][2][2]int)
	for _, word := range words {
		loc, found := grid.Search(word)
		result[word] = loc
		if !found {
			err = fmt.Errorf("word %s not found", word)
		}
	}

	return result, err
}

type Grid struct {
	grid [][]rune
}

func NewGrid(puzzle []string) *Grid {
	grid := make([][]rune, len(puzzle))
	for i, row := range puzzle {
		grid[i] = []rune(row)
	}
	return &Grid{grid: grid}
}

func (g Grid) String() string {
	return fmt.Sprintf("%v", g.grid)
}

func (g Grid) Search(word string) (result [2][2]int, found bool) {
	for row_i, row := range g.grid {
		for col_i := range row {
			for _, direction := range DIRECTIONS {
				if isMatch(g, word, row_i, col_i, direction) {
					startLocation := [2]int{col_i, row_i}
					endLocation := endLocation(startLocation, direction, word)
					return [2][2]int{startLocation, endLocation}, true
				}
			}
		}
	}
	return NOT_FOUND_LINE, false
}

func (g Grid) SafeGet(row int, col int) (rune, bool) {
	if row < 0 || row >= len(g.grid) {
		return ' ', false
	}
	if col < 0 || col >= len(g.grid[row]) {
		return ' ', false
	}
	return g.grid[row][col], true
}

func isMatch(grid Grid, word string, row int, col int, direction string) bool {
	gridWord := []rune{}
	for i := 0; i < len(word); i++ {
		switch direction {
		case NORTH:
			if char, ok := grid.SafeGet(row-i, col); ok {
				gridWord = append(gridWord, char)
			}
		case EAST:
			if char, ok := grid.SafeGet(row, col+i); ok {
				gridWord = append(gridWord, char)
			}
		case SOUTH:
			if char, ok := grid.SafeGet(row+i, col); ok {
				gridWord = append(gridWord, char)
			}
		case WEST:
			if char, ok := grid.SafeGet(row, col-i); ok {
				gridWord = append(gridWord, char)
			}
		case NORTH_EAST:
			if char, ok := grid.SafeGet(row-i, col+i); ok {
				gridWord = append(gridWord, char)
			}
		case NORTH_WEST:
			if char, ok := grid.SafeGet(row-i, col-i); ok {
				gridWord = append(gridWord, char)
			}
		case SOUTH_EAST:
			if char, ok := grid.SafeGet(row+i, col+i); ok {
				gridWord = append(gridWord, char)
			}
		case SOUTH_WEST:
			if char, ok := grid.SafeGet(row+i, col-i); ok {
				gridWord = append(gridWord, char)
			}
		}
	}
	return string(gridWord) == word
}

func endLocation(startLoc [2]int, direction string, word string) (endLoc [2]int) {
	col, row := startLoc[0], startLoc[1]
	delta := len(word) - 1
	switch direction {
	case NORTH:
		return [2]int{col, row - delta}
	case EAST:
		return [2]int{col + delta, row}
	case SOUTH:
		return [2]int{col, row + delta}
	case WEST:
		return [2]int{col - delta, row}
	case NORTH_EAST:
		return [2]int{col + delta, row - delta}
	case NORTH_WEST:
		return [2]int{col - delta, row - delta}
	case SOUTH_EAST:
		return [2]int{col + delta, row + delta}
	case SOUTH_WEST:
		return [2]int{col - delta, row + delta}
	default:
		panic(fmt.Sprintf("invalid direction %v", direction))
	}
}
