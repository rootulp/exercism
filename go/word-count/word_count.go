package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) (result Frequency) {
	result = Frequency{}
	splitter := func(c rune) bool {
		return unicode.IsSpace(c) || (unicode.IsPunct(c) && c != '\'')
	}
	fields := strings.FieldsFunc(phrase, splitter)
	for _, word := range fields {
		if doesContainsLetterOrNumber(word) {
			w := strings.Trim(strings.ToLower(word), "'")
			result[strings.ToLower(w)] += 1
		}
	}
	return result
}

func doesContainsLetterOrNumber(word string) bool {
	containsLetterOrNumber := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsNumber(c)
	}
	return strings.IndexFunc(word, containsLetterOrNumber) != -1
}
