package minesweeper

import (
	"strconv"
	"strings"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	if noRows(board) || noColumns(board) {
		return board
	}

	game := NewGame(board)
	game.solve()
	return game.AnnotatedBoard()
}

func noRows(board []string) bool {
	return len(board) == 0
}

func noColumns(board []string) bool {
	return len(board[0]) == 0
}

type Game struct {
	grid          [][]string
	gridAnnotated [][]string
}

func NewGame(board []string) Game {
	grid := [][]string{}
	gridAnnotated := [][]string{}

	for _, row := range board {
		grid = append(grid, strings.Split(row, ""))
		gridAnnotated = append(gridAnnotated, strings.Split(row, ""))
	}

	return Game{
		grid:          grid,
		gridAnnotated: gridAnnotated,
	}
}

func (g Game) AnnotatedBoard() (board []string) {
	for _, row := range g.gridAnnotated {
		board = append(board, strings.Join(row, ""))
	}
	return board
}

func (g Game) solve() {
	for row_i, row := range g.grid {
		for col_i, cell := range row {
			if isMine(cell) {
				continue
			}
			mines := g.neighborMines(row_i, col_i)
			if mines == 0 {
				continue
			}
			g.gridAnnotated[row_i][col_i] = strconv.Itoa(mines)
		}
	}
}

func (g Game) neighborMines(row_i, col_i int) int {
	neighbors := g.getNeighbors(row_i, col_i)
	mines := 0
	for _, neighbor := range neighbors {
		if isMine(neighbor) {
			mines++
		}
	}
	return mines
}

func (g Game) getNeighbors(row_i, col_i int) (neighbors []string) {
	for row := row_i - 1; row <= row_i+1; row++ {
		for col := col_i - 1; col <= col_i+1; col++ {
			if row < 0 || col < 0 || row >= len(g.grid) || col >= len(g.grid[row]) {
				continue
			}
			if row == row_i && col == col_i {
				continue
			}
			neighbors = append(neighbors, g.grid[row][col])
		}
	}
	return neighbors
}

func isMine(cell string) bool {
	return cell == "*"
}
