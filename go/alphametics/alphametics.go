// Package alphametics provides functionality for solving alphametics puzzles.
package alphametics

import (
	"errors"
	"math"
	"sort"
	"strings"
)

type addend struct {
	letters []rune
}

type equation struct {
	usedLetters []rune
	lhs         []addend
	rhs         addend
}

// Solve solves the given input puzzle.
func Solve(puzzle string) (map[string]int, error) {
	eq := parsePuzzle(puzzle)
	result := solvePuzzle(eq)

	if result == nil {
		return nil, errors.New("no possible solution")
	}

	return result, nil
}

func parsePuzzle(puzzle string) equation {
	lhs := make([]addend, 0)
	var rhs addend
	usedLetters := make([]rune, 0, 10)
	isRHS := false

	for _, item := range strings.Fields(puzzle) {
		if item == "==" {
			isRHS = true
			continue
		} else if item == "+" {
			continue
		} else if isRHS {
			rhs = addend{letters: []rune(item)}
		} else {
			lhs = append(lhs, addend{letters: []rune(item)})
		}

		for _, v := range item {
			if !isLetterUsed(usedLetters, v) {
				usedLetters = append(usedLetters, v)
			}
		}
	}

	sort.Slice(usedLetters, func(i, j int) bool { return usedLetters[i] < usedLetters[j] })
	return equation{usedLetters: usedLetters, lhs: lhs, rhs: rhs}
}

func solvePuzzle(eq equation) map[string]int {
	result := map[string]int{}
	usedNumbers := make([]int, 0, 10)
	isSolved := false

outer:
	for !isSolved {
		isSolved, usedNumbers = solveEquation(eq, usedNumbers)

		if !isSolved {
			for usedNumbers[len(usedNumbers)-1] == getMaxNumberCanBeUsed(usedNumbers) {
				if len(usedNumbers) == 1 {
					break outer // no more possible solution
				}

				usedNumbers = usedNumbers[:len(usedNumbers)-1]
			}

			for i := usedNumbers[len(usedNumbers)-1] + 1; i < 10; i++ {
				if isNumberUsed(usedNumbers, i) {
					continue
				} else {
					usedNumbers[len(usedNumbers)-1] = i
					break
				}
			}
		}
	}

	if isSolved {
		for i := 0; i < len(usedNumbers); i++ {
			result[string(eq.usedLetters[i])] = usedNumbers[i]
		}

		return result
	}

	return nil
}

func solveEquation(eq equation, usedNumbers []int) (bool, []int) {
	for i := 0; i < 10; i++ {
		if isNumberUsed(usedNumbers, i) {
			continue
		}

		usedNumbers = append(usedNumbers, i)

		if len(usedNumbers) == len(eq.usedLetters) {
			if isEquationTrue(eq, usedNumbers) {
				return true, usedNumbers
			}

			usedNumbers = usedNumbers[:len(usedNumbers)-1]
		}
	}

	return false, usedNumbers
}

func isEquationTrue(eq equation, usedNumbers []int) bool {
	var result bool
	var lhs, lhsSum, rhs int

	for i := 0; i < len(eq.lhs); i++ {
		if result, lhs = getSum(eq.lhs[i], eq.usedLetters, usedNumbers); !result {
			return result
		}

		lhsSum += lhs
	}

	if result, rhs = getSum(eq.rhs, eq.usedLetters, usedNumbers); !result {
		return result
	}

	return lhsSum == rhs
}

func isNumberUsed(usedNumbers []int, number int) bool {
	for i := 0; i < len(usedNumbers); i++ {
		if usedNumbers[i] == number {
			return true
		}
	}

	return false
}

func isLetterUsed(usedLetters []rune, letter rune) bool {
	for i := 0; i < len(usedLetters); i++ {
		if usedLetters[i] == letter {
			return true
		}
	}

	return false
}

func getMaxNumberCanBeUsed(usedNumbers []int) int {
	for i := 9; i >= 0; i-- {
		if !isNumberUsed(usedNumbers[:len(usedNumbers)-1], i) {
			return i
		}
	}

	return 0
}

func getSum(ad addend, usedLetters []rune, usedNumbers []int) (bool, int) {
	var sum int
	numOfDigits := len(ad.letters)

	for i := 0; i < numOfDigits; i++ {
		num := getNumberFromUsedNumbers(ad.letters[i], usedLetters, usedNumbers)

		if num == 0 && i == 0 && numOfDigits > 1 {
			return false, -1 // leading zero in multi-digit number
		}

		sum += num * int(math.Pow(10, float64(numOfDigits-1-i)))
	}

	return true, sum
}

func getNumberFromUsedNumbers(letter rune, usedLetters []rune, usedNumbers []int) int {
	for i := 0; i < len(usedLetters); i++ {
		if letter == usedLetters[i] {
			return usedNumbers[i]
		}
	}

	return -1
}
