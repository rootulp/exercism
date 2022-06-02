package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	grid [][]int
}

func (m *Matrix) String() string {
	return fmt.Sprintf("grid: %v", m.grid)
}

func (m *Matrix) isSaddle(x int, y int) bool {
	return m.isMaxInRow(x, y) && m.isMinInCol(x, y)
}

func (m *Matrix) isMaxInRow(x int, y int) bool {
	val := m.grid[y][x]

	for _, i := range m.grid[y] {
		if val < i {
			return false
		}
	}
	return true
}

func (m *Matrix) isMinInCol(x int, y int) bool {
	val := m.grid[y][x]

	for _, i := range m.column(x) {
		if val > i {
			return false
		}
	}
	return true
}

func (m *Matrix) column(x int) (col []int) {
	for _, row := range m.grid {
		col = append(col, row[x])
	}
	return col
}

type Pair struct {
	row int
	col int
}

func New(s string) (m *Matrix, err error) {
	grid := [][]int{}

	for _, line := range strings.Split(s, "\n") {
		row := []int{}
		for _, c := range strings.Fields(line) {
			i, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		grid = append(grid, row)
	}
	return &Matrix{grid}, nil
}

func (m *Matrix) Saddle() (pairs []Pair) {
	pairs = []Pair{}

	for y, row := range m.grid {
		for x := range row {
			if m.isSaddle(x, y) {
				pair := Pair{row: y, col: x}
				pairs = append(pairs, pair)
			}
		}
	}
	return pairs
}
