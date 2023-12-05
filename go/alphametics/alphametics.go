package alphametics

import (
	"errors"
	"fmt"
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

	return map[string]int{}, nil
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
