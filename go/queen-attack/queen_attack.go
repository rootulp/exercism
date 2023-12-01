package queenattack

import (
	"errors"
	"fmt"
	"strconv"
)

type location struct {
	row    int
	column string
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	white, err := parse(whitePosition)
	if err != nil {
		return false, err
	}
	black, err := parse(blackPosition)
	if err != nil {
		return false, err
	}
	if white == black {
		return false, errors.New("same square")
	}
	fmt.Printf("white: %v, black: %v\n", white, black)
	return false, nil
}

func parse(position string) (location, error) {
	if position == "" {
		return location{}, errors.New("empty position")
	}
	if len(position) != 2 {
		return location{}, errors.New("invalid position")
	}

	column := string(position[0])
	row, err := strconv.Atoi(position[1:])

	if err != nil {
		return location{}, errors.New("invalid position")
	}

	if !isValidRow(row) {
		return location{}, errors.New("invalid row")
	}

	if !isValidColumn(column) {
		return location{}, errors.New("invalid column")
	}

	return location{row: row, column: column}, nil
}

func isValidRow(row int) bool {
	return row >= 1 && row <= 8
}

func isValidColumn(column string) bool {
	return column >= "a" && column <= "h"
}
