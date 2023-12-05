package alphametics

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func Solve(input string) (map[string]int, error) {
	parts := strings.Split(input, "==")
	if len(parts) != 2 {
		return nil, errors.New("invalid input")
	}
	parts = trim(parts)
	left := parts[0]
	sum := parts[1]
	addends := strings.Split(left, "+")
	addends = trim(addends)

	fmt.Printf("left: %s\n", left)
	fmt.Printf("addends: %s\n", addends)
	fmt.Printf("sum: %s\n", sum)
	words := append([]string{sum}, addends...)

	letters := getUniqueLetters(words)
	fmt.Printf("letters %v\n", letters)

	for {
		letterToNumber := assignRandomValues(letters)
		fmt.Printf("letterToNumber %v\n", letterToNumber)

		evaluated, err := evaluate(sum, addends, letterToNumber)
		if err != nil {
			return nil, err
		}
		if evaluated {
			return letterToNumber, nil
		}
	}
}

func evaluate(sum string, addends []string, letterToNumber map[string]int) (bool, error) {
	result := 0
	for _, addend := range addends {
		translated, err := translate(letterToNumber, addend)
		if err != nil {
			return false, err
		}
		result += translated
	}
	realSum, err := translate(letterToNumber, sum)
	if err != nil {
		return false, err
	}
	return result == realSum, nil
}

func translate(letterToNumber map[string]int, word string) (int, error) {
	translated := ""
	for _, letter := range word {
		value := letterToNumber[string(letter)]
		translated += strconv.Itoa(value)
	}
	result, err := strconv.Atoi(translated)
	if err != nil {
		return 0, err
	}
	return result, nil
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

func getUniqueLetters(words []string) map[string]bool {
	result := map[string]bool{}
	for _, word := range words {
		for _, letter := range word {
			result[string(letter)] = true
		}
	}
	return result
}

// trim removes leading and trailing spaces from each string in x
func trim(x []string) []string {
	for i, addend := range x {
		x[i] = strings.Trim(addend, " ")
	}
	return x
}
