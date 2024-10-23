package alphametics

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (map[string]int, error) {
	equation, err := parse(input)
	if err != nil {
		return nil, err
	}
	if equation.onlyTwoUniqueLetters() {
		return nil, fmt.Errorf("solution must have unique value for each letter")
	}
	if equation.addendIsLongerThanSum() {
		return nil, fmt.Errorf("leading zero solution is invalid")
	}

	letters := equation.uniqueLetters()
	usedNumbers := make([]bool, 10)
	letterToNumber := make(map[string]int)

	if solveBacktrack(equation, letters, letterToNumber, usedNumbers, 0) {
		return letterToNumber, nil
	}
	return nil, errors.New("no solution found")
}

func solveBacktrack(equation equation, letters []string, letterToNumber map[string]int, usedNumbers []bool, index int) bool {
	if index == len(letters) {
		return equation.evaluate(letterToNumber)
	}

	for num := 0; num <= 9; num++ {
		if !usedNumbers[num] {
			letterToNumber[letters[index]] = num
			usedNumbers[num] = true
			if !equation.isLeadingZero(letterToNumber) && solveBacktrack(equation, letters, letterToNumber, usedNumbers, index+1) {
				return true
			}

			usedNumbers[num] = false
			delete(letterToNumber, letters[index])
		}
	}
	return false
}

type equation struct {
	addends []string
	sum     string
}

// parse returns an equation based on input
func parse(input string) (equation, error) {
	parts := trim(strings.Split(input, "=="))
	if len(parts) != 2 {
		return equation{}, errors.New("invalid input")
	}
	left, sum := parts[0], parts[1]

	addends := trim(strings.Split(left, "+"))
	return equation{addends, sum}, nil
}

// words returns all words in the equation
func (e equation) words() []string {
	return append([]string{e.sum}, e.addends...)
}

// uniqueLeters returns the unique letters in the equation.
func (e equation) uniqueLetters() []string {
	letters := map[string]bool{}
	for _, word := range e.words() {
		for _, letter := range word {
			letters[string(letter)] = true
		}
	}
	return getKeys(letters)
}

// getKeys returns all keys in a map
func getKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// evaluate returns true if the equation is satisfied based on letterToNumber
func (e equation) evaluate(letterToNumber map[string]int) bool {
	return e.evaluateExpression(letterToNumber) == translate(letterToNumber, e.sum)
}

// evaluateExpression returns the result of the sum of all addends based on letterToNumber
func (e equation) evaluateExpression(letterToNumber map[string]int) (result int) {
	for _, addend := range e.addends {
		translated := translate(letterToNumber, addend)
		result += translated
	}
	return result
}

// isLeadingZero returns true if any word has a leading zero based on letterToNumber.
func (e equation) isLeadingZero(letterToNumber map[string]int) bool {
	for _, word := range e.words() {
		if len(word) > 1 {
			firstLetter := string(word[0])
			if value, exists := letterToNumber[firstLetter]; exists && value == 0 {
				return true
			}
		}
	}
	return false
}

// onlyTwoUniqueLetters returns true if there are only two unique letters in the equation.
func (e equation) onlyTwoUniqueLetters() bool {
	return len(e.uniqueLetters()) == 2
}

// addendIsLongerThanSum returns true if any addend is longer than sum
func (e equation) addendIsLongerThanSum() bool {
	for _, addend := range e.addends {
		if len(addend) > len(e.sum) {
			return true
		}
	}
	return false
}

// translate converts word to a number based on letterToNumber
func translate(letterToNumber map[string]int, word string) int {
	numbers := ""
	for _, letter := range word {
		value := letterToNumber[string(letter)]
		numbers += strconv.Itoa(value)
	}
	value, err := strconv.Atoi(numbers)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to int", numbers))
	}
	return value
}

// trim removes leading and trailing spaces from each word in words
func trim(words []string) []string {
	for i, word := range words {
		words[i] = strings.Trim(word, " ")
	}
	return words
}
