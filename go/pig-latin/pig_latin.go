package piglatin

import (
	"fmt"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}
var consonantSounds = []string{"squ", "thr", "th", "sch", "qu", "ch", "rh", "p", "k", "x", "q", "y", "m", "f", "r"}

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	translated := []string{}
	for _, word := range words {
		translated = append(translated, translate(word))
	}
	return strings.Join(translated, " ")
}

func translate(word string) string {
	if startsWithVowel(word) ||
		strings.HasPrefix(word, "xr") ||
		strings.HasPrefix(word, "yt") {
		return handleVowel(word)
	}
	if startsWithConsonantSound(word) {
		return handleConsonantSound(word)
	}
	panic(fmt.Sprintf("unhandled word %v", word))
}

func startsWithVowel(sentence string) bool {
	for _, vowel := range vowels {
		if strings.HasPrefix(sentence, vowel) {
			return true
		}
	}
	return false
}

func startsWithConsonantSound(sentence string) bool {
	for _, cs := range consonantSounds {
		if strings.HasPrefix(sentence, cs) {
			return true
		}
	}
	return false
}

func handleConsonantSound(sentence string) string {
	for _, cs := range consonantSounds {
		if strings.HasPrefix(sentence, cs) {
			return strings.TrimPrefix(sentence, cs) + cs + "ay"
		}
	}
	panic(fmt.Sprintf("could not find consonant sound prefix for sentence %v", sentence))
}

func handleVowel(sentence string) string {
	return sentence + "ay"
}
