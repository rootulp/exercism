package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	grid [][]int
}

func New(s string) (*Matrix, error) {
	if isUnevenRows(s) {
		return &Matrix{}, fmt.Errorf("uneven rows for %s", s)
	}
	rows := strings.Split(s, "\n")
	grid := [][]int{}
	for _, row := range rows {
		r := []int{}
		cols := strings.Split(strings.Trim(row, " "), " ")
		for _, col := range cols {
			val, err := strconv.Atoi(col)
			if err != nil {
				return &Matrix{}, err
			}
			r = append(r, val)
		}
		grid = append(grid, r)
	}
	return &Matrix{
		grid: grid,
	}, nil
}

func (m Matrix) Cols() [][]int {
	result := initializeGrid(len(m.grid[0]), len(m.grid))
	for x, row := range m.grid {
		for y, val := range row {
			result[y][x] = val
		}
	}
	return result
}

func (m Matrix) Rows() [][]int {
	// return m.grid
	duplicate := make([][]int, len(m.grid))
	for i := range m.grid {
		duplicate[i] = make([]int, len(m.grid[i]))
		copy(duplicate[i], m.grid[i])
	}
	return duplicate
}

func (m Matrix) Set(row int, column int, value int) bool {
	if row >= len(m.grid) || row < 0 {
		return false
	}
	if column >= len(m.grid[0]) || column < 0 {
		return false
	}
	m.grid[row][column] = value
	return true
}

func isUnevenRows(s string) bool {
	rowLengths := map[int]bool{}
	rows := strings.Split(s, "\n")
	for _, row := range rows {
		cols := strings.Split(strings.Trim(row, " "), " ")
		rowLengths[len(cols)] = true
	}
	return len(rowLengths) != 1
}

func initializeGrid(dy int, dx int) (grid [][]int) {
	result := make([][]int, dy)
	for i := range result {
		result[i] = make([]int, dx)
	}
	return result
}
