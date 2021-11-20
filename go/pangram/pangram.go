package pangram

import "strings"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(input string) bool {
	set := getSetOfUsedRunes(strings.ToLower(input))
	alphabetSet := getSetOfUsedRunes(alphabet)
	for r, v := range alphabetSet {
		if set[r] != v {
			return false
		}
	}
	return true
}

func getSetOfUsedRunes(input string) map[rune]bool {
	set := map[rune]bool{}
	for _, r := range input {
		set[r] = true
	}
	return set
}
