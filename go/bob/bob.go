package bob

import (
	"strings"
	"unicode"
)

// Hey returns a string with Bob's response.
func Hey(remark string) string {
	if isEmpty(remark) {
		return "Fine. Be that way!"
	} else if isYellingQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		return "Sure."
	} else if isYelling(remark) {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

func removeNonLetters(remark string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, remark)
}

func isYellingQuestion(remark string) bool {
	return isYelling(remark) && isQuestion(remark)
}

func isYelling(s string) bool {
	if removeNonLetters(s) == s {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
