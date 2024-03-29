package stateoftictactoe

import (
	"errors"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	game := NewGame(board)
	return game.getState()
}

type game struct {
	grid [][]string
}

func NewGame(board []string) *game {
	grid := [][]string{}
	for _, row := range board {
		grid = append(grid, strings.Split(row, ""))
	}
	return &game{grid: grid}
}

func (g game) getState() (State, error) {
	if err := g.isInvalidBoard(); err != nil {
		return "", err
	}
	if isWin := g.isWin("X") || g.isWin("O"); isWin {
		return Win, nil
	}
	if isDraw := g.isDraw(); isDraw {
		return Draw, nil
	}
	return Ongoing, nil
}

func (g game) isWin(player string) bool {
	return g.isWinByRow(player) || g.isWinByColumn(player) || g.isWinByDiagonal(player)
}

func (g game) isWinByRow(player string) bool {
	for _, row := range g.grid {
		if allCellsOccupied(player, row...) {
			return true
		}
	}
	return false
}

func (g game) isWinByColumn(player string) bool {
	for i := 0; i < 3; i++ {
		if allCellsOccupied(player, g.grid[0][i], g.grid[1][i], g.grid[2][i]) {
			return true
		}
	}
	return false
}

func (g game) isWinByDiagonal(player string) bool {
	if allCellsOccupied(player, g.grid[0][0], g.grid[1][1], g.grid[2][2]) {
		return true
	}
	if allCellsOccupied(player, g.grid[0][2], g.grid[1][1], g.grid[2][0]) {
		return true
	}
	return false
}

// allCellsOccupied returns true if all cells are occupied by player.
func allCellsOccupied(player string, cells ...string) bool {
	for _, cell := range cells {
		if cell != player {
			return false
		}
	}
	return true
}

func (g game) isDraw() bool {
	isWin := g.isWin("X") || g.isWin("O")
	return g.isOver() && !isWin
}

func (g game) isOver() bool {
	return g.getCount("X")+g.getCount("O") == 9
}

func (g game) isInvalidBoard() error {
	movesX := g.getCount("X")
	movesO := g.getCount("O")
	if movesX > movesO+1 {
		return errors.New("invalid board")
	}
	if movesO > movesX {
		return errors.New("invalid board")
	}
	if g.isWin("X") && g.isWin("O") {
		return errors.New("invalid board")
	}
	return nil
}

func (g game) getCount(player string) (count int) {
	for _, row := range g.grid {
		for _, cell := range row {
			if cell == player {
				count++
			}
		}
	}
	return count
}
