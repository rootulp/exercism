package queenattack

import (
	"errors"
	"strconv"
)

var fileToColumn = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
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

func parse(position string) (location, error) {
	if len(position) != 2 {
		return location{}, errors.New("invalid position")
	}

	file := string(position[0])
	row, err := strconv.Atoi(position[1:])
	if err != nil || !isValidRow(row) || !isValidFile(file) {
		return location{}, errors.New("invalid position")
	}
	column := fileToColumn[file]

	return location{row: row, column: column}, nil
}

func isValidRow(row int) bool {
	return row >= 1 && row <= 8
}

func isValidFile(column string) bool {
	return column >= "a" && column <= "h"
}

type location struct {
	row    int
	column int
}

func (l location) canAttack(other location) bool {
	return l.isSameRow(other) || l.isSameColumn(other) || l.isDiagonal(other)
}

func (l location) isSameRow(other location) bool {
	return l.row == other.row
}

func (l location) isSameColumn(other location) bool {
	return l.column == other.column
}

func (l location) isDiagonal(other location) bool {
	return abs(l.row-other.row) == abs(l.column-other.column)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
