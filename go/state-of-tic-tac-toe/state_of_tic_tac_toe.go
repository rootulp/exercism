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
	movesX := g.getCount("X")
	movesO := g.getCount("O")
	if movesX > movesO+1 {
		return "", errors.New("invalid board")
	}
	return Ongoing, nil
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
