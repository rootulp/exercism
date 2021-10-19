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
	return [][]int{}
}

func (m Matrix) Rows() [][]int {
	return m.grid
}

func (m Matrix) Set(row int, column int, value int) bool {
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
