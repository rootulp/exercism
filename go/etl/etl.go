package etl

import "strings"

func Transform(in map[int][]string) (result map[string]int) {
	result = map[string]int{}
	for score, letters := range in {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}
