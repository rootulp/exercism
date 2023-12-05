package alphametics

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func Solve(input string) (map[string]int, error) {
	equation, err := parse(input)
	if err != nil {
		return nil, err
	}
	if equation.nonUniqueValues() {
		return nil, fmt.Errorf("solution must have unique value for each letter")
	}

	letters := equation.uniqueLetters()
	for {
		letterToNumber := assignRandomValues(letters)
		if equation.isLeadingZero(letterToNumber) {
			continue
		}
		if equation.evaluate(letterToNumber) {
			return letterToNumber, nil
		}
	}
}

type equation struct {
	addends []string
	sum     string
}

func parse(input string) (equation, error) {
	parts := trim(strings.Split(input, "=="))
	if len(parts) != 2 {
		return equation{}, errors.New("invalid input")
	}
	left, sum := parts[0], parts[1]

	addends := trim(strings.Split(left, "+"))
	return equation{addends, sum}, nil
}

func (e equation) words() []string {
	return append([]string{e.sum}, e.addends...)
}

func (e equation) uniqueLetters() map[string]bool {
	result := map[string]bool{}
	for _, word := range e.words() {
		for _, letter := range word {
			result[string(letter)] = true
		}
	}
	return result
}

func (e equation) evaluate(letterToNumber map[string]int) bool {
	return e.evaluateExpression(letterToNumber) == translate(letterToNumber, e.sum)
}

func (e equation) evaluateExpression(letterToNumber map[string]int) (result int) {
	for _, addend := range e.addends {
		translated := translate(letterToNumber, addend)
		result += translated
	}
	return result
}

func (e equation) isLeadingZero(letterToNumber map[string]int) bool {
	for _, word := range e.words() {
		if len(word) > 1 && letterToNumber[string(word[0])] == 0 {
			return true
		}
	}
	return false
}

func (e equation) nonUniqueValues() bool {
	return len(e.uniqueLetters()) == 2
}

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

func assignRandomValues(letters map[string]bool) map[string]int {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := map[string]int{}
	for letter := range letters {
		value, remainingElements := popRandom(values)
		values = remainingElements
		result[letter] = value
	}
	return result
}

func popRandom(x []int) (int, []int) {
	index := rand.Intn(len(x))
	element := x[index]
	remaining := append(x[:index], x[index+1:]...)
	return element, remaining
}

// trim removes leading and trailing spaces from each word in words
func trim(words []string) []string {
	for i, word := range words {
		words[i] = strings.Trim(word, " ")
	}
	return words
}
