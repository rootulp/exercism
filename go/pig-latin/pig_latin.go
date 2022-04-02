package piglatin

import (
	"fmt"
	"strings"
)

var consonants = []string{"p", "k"}

func Sentence(sentence string) string {
	if startsWithVowel(sentence) ||
		strings.HasPrefix(sentence, "xr") ||
		strings.HasPrefix(sentence, "yt") {
		return sentence + "ay"
	}
	if startsWithConsonant(sentence) {
		return handleConsonant(sentence)
	}

	return sentence
}

func startsWithVowel(sentence string) bool {
	return strings.HasPrefix(sentence, "a") ||
		strings.HasPrefix(sentence, "e") ||
		strings.HasPrefix(sentence, "i") ||
		strings.HasPrefix(sentence, "o") ||
		strings.HasPrefix(sentence, "u")
}

func startsWithConsonant(sentence string) bool {
	for _, consonant := range consonants {
		if strings.HasPrefix(sentence, consonant) {
			return true
		}
	}
	return false
}

func handleConsonant(sentence string) string {
	for _, consonant := range consonants {
		if strings.HasPrefix(sentence, consonant) {
			return strings.TrimPrefix(sentence, consonant) + consonant + "ay"
		}
	}
	panic(fmt.Sprintf("could not find consonant prefix for sentence %v", sentence))
}
