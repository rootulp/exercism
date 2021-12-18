package allergies

import (
	"sort"
)

var scoreToAllergen = map[int]string{
	128: "cats",
	64:  "pollen",
	32:  "chocolate",
	16:  "tomatoes",
	8:   "strawberries",
	4:   "shellfish",
	2:   "peanuts",
	1:   "eggs",
}

func Allergies(score uint) (allergies []string) {
	if score >= 256 {
		score = score % 256
	}

	for _, key := range descendingKeys(scoreToAllergen) {
		allergen := scoreToAllergen[key]
		if score/uint(key) == 1 {
			allergies = append(allergies, allergen)
			score = score % uint(key)
		}
	}
	return allergies
}

func AllergicTo(score uint, allergen string) bool {
	allergies := Allergies(score)
	for _, candidate := range allergies {
		if candidate == allergen {
			return true
		}
	}
	return false
}

func descendingKeys(m map[int]string) (keys []int) {
	for k := range scoreToAllergen {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	sort.Slice(keys, func(a, b int) bool {
		return keys[b] < keys[a]
	})
	return keys
}
