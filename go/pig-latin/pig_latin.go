package piglatin

import "strings"

func Sentence(sentence string) string {
	if startsWithVowel(sentence) ||
		strings.HasPrefix(sentence, "xr") ||
		strings.HasPrefix(sentence, "yt") {
		return sentence + "ay"
	}
	if strings.HasPrefix(sentence, "p") {
		return strings.TrimPrefix(sentence, "p") + "p" + "ay"
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
