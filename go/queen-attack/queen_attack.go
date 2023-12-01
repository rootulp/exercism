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

func parse(position string) (coordinate, error) {
	if len(position) != 2 {
		return coordinate{}, errors.New("invalid position")
	}

	file := string(position[0])
	row, err := strconv.Atoi(position[1:])
	if err != nil || !isValidRow(row) || !isValidFile(file) {
		return coordinate{}, errors.New("invalid position")
	}
	column := fileToColumn[file]

	return coordinate{row: row, column: column}, nil
}

func isValidRow(row int) bool {
	return row >= 1 && row <= 8
}

func isValidFile(file string) bool {
	return file >= "a" && file <= "h"
}

type coordinate struct {
	row    int
	column int
}

func (c coordinate) canAttack(other coordinate) bool {
	return c.isSameRow(other) || c.isSameColumn(other) || c.isDiagonal(other)
}

func (c coordinate) isSameRow(other coordinate) bool {
	return c.row == other.row
}

func (c coordinate) isSameColumn(other coordinate) bool {
	return c.column == other.column
}

func (c coordinate) isDiagonal(other coordinate) bool {
	return abs(c.row-other.row) == abs(c.column-other.column)
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
