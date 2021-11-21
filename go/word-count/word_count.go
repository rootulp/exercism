package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) (result Frequency) {
	result = Frequency{}
	words := splitStringOnWhitespaceAndPunctuation(phrase)
	for _, word := range words {
		if doesContainsLetterOrNumber(word) {
			w := strings.Trim(strings.ToLower(word), "'")
			result[strings.ToLower(w)] += 1
		}
	}
	return result
}

func splitStringOnWhitespaceAndPunctuation(phrase string) []string {
	splitter := func(c rune) bool {
		return unicode.IsSpace(c) || (unicode.IsPunct(c) && c != '\'')
	}
	return strings.FieldsFunc(phrase, splitter)
}

func doesContainsLetterOrNumber(word string) bool {
	containsLetterOrNumber := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsNumber(c)
	}
	return strings.IndexFunc(word, containsLetterOrNumber) != -1
}
