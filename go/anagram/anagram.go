package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) (anagrams []string) {
	occurences := getOccurences(strings.ToLower(subject))
	for _, c := range candidates {
		candidate := getOccurences(strings.ToLower(c))
		if reflect.DeepEqual(occurences, candidate) && strings.ToLower(subject) != strings.ToLower(c) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func getOccurences(input string) (occurences map[rune]int) {
	occurences = map[rune]int{}
	for _, r := range input {
		occurences[r] += 1
	}
	return occurences
}
