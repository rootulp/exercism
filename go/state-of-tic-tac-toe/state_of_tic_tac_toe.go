package stateoftictactoe

import (
	"errors"
	"fmt"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	game, err := parse(board)
	if err != nil {
		return "", err
	}
	return game.getState()
}

type game struct {
	grid [][]string
}

func parse(board []string) (*game, error) {
	grid := [][]string{}
	for _, row := range board {
		grid = append(grid, strings.Split(row, ""))
	}
	fmt.Printf("%#v\n", grid)
	return &game{grid: grid}, nil
}

func (g game) getState() (State, error) {
	if err := g.isInvalidBoard(); err != nil {
		return "", err
	}
	if isWin := g.isWin(); isWin {
		return Win, nil
	}
	if isDraw := g.isDraw(); isDraw {
		return Draw, nil
	}
	return Ongoing, nil
}

func (g game) isWin() bool {
	return g.isWinByRow() || g.isWinByColumn() || g.isWinByDiagonal()
}

func (g game) isWinByRow() bool {
	for _, row := range g.grid {
		if row[0] == row[1] && row[1] == row[2] && row[0] != " " {
			return true
		}
	}
	return false
}

func (g game) isWinByColumn() bool {
	for i := 0; i < 3; i++ {
		if g.grid[0][i] == g.grid[1][i] && g.grid[1][i] == g.grid[2][i] && g.grid[0][i] != " " {
			return true
		}
	}
	return false
}

func (g game) isWinByDiagonal() bool {
	if g.grid[0][0] == g.grid[1][1] && g.grid[1][1] == g.grid[2][2] && g.grid[0][0] != " " {
		return true
	}
	if g.grid[0][2] == g.grid[1][1] && g.grid[1][1] == g.grid[2][0] && g.grid[0][2] != " " {
		return true
	}
	return false
}

func (g game) isDraw() bool {
	return g.getCount("X")+g.getCount("O") == 9 && !g.isWin()
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
