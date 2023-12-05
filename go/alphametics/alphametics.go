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

	letters := equation.uniqueLetters()
	for {
		letterToNumber := assignRandomValues(letters)
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

func translate(letterToNumber map[string]int, word string) int {
	translated := ""
	for _, letter := range word {
		value := letterToNumber[string(letter)]
		translated += strconv.Itoa(value)
	}
	result, err := strconv.Atoi(translated)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to int", translated))
	}
	return result
}

func assignRandomValues(letters map[string]bool) map[string]int {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := map[string]int{}
	for letter := range letters {
		value, remainingElements := popRandomElement(values)
		values = remainingElements
		result[letter] = value
	}
	return result
}

func popRandomElement(x []int) (int, []int) {
	index := rand.Intn(len(x))
	result := x[index]
	x = append(x[:index], x[index+1:]...)
	return result, x
}

// trim removes leading and trailing spaces from each string in x
func trim(x []string) []string {
	for i, addend := range x {
		x[i] = strings.Trim(addend, " ")
	}
	return x
}
