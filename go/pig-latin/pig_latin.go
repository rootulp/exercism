package piglatin

import (
	"fmt"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}
var consonantSounds = []string{"squ", "thr", "th", "sch", "qu", "ch", "rh", "p", "k", "x", "q", "y", "m"}

func Sentence(sentence string) string {
	if startsWithVowel(sentence) ||
		strings.HasPrefix(sentence, "xr") ||
		strings.HasPrefix(sentence, "yt") {
		return handleVowel(sentence)
	}
	if startsWithConsonantSound(sentence) {
		return handleConsonantSound(sentence)
	}

	return sentence
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
