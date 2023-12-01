package queenattack

import (
	"errors"
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
	return white.canAttack(black), nil
}

func (l location) canAttack(other location) bool {
	return l.isSameRow(other) || l.isSameColumn(other)
}

func (l location) isSameRow(other location) bool {
	return l.row == other.row
}

func (l location) isSameColumn(other location) bool {
	return l.column == other.column
}

// func (l location) isDiagonal(other location) bool {
// 	return abs(l.row-other.row) == abs(int(l.column[0])-int(other.column[0]))
// }

func parse(position string) (location, error) {
	if len(position) != 2 {
		return location{}, errors.New("invalid position")
	}

	column := string(position[0])
	row, err := strconv.Atoi(position[1:])

	if err != nil || !isValidRow(row) || !isValidColumn(column) {
		return location{}, errors.New("invalid position")
	}

	return location{row: row, column: column}, nil
}

func isValidRow(row int) bool {
	return row >= 1 && row <= 8
}

func isValidColumn(column string) bool {
	return column >= "a" && column <= "h"
}
