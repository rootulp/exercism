package wordsearch

import "fmt"

func Solve(words []string, puzzle []string) (result map[string][2][2]int, err error) {
	result = make(map[string][2][2]int)
	for _, word := range words {
		result[word] = [2][2]int{{-1, -1}, {-1, -1}}
	}
	for k, v := range result {
		if v == [2][2]int{{-1, -1}, {-1, -1}} {
			err = fmt.Errorf("word %s not found", k)
		}
	}
	return result, err
}
