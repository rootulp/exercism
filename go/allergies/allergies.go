package allergies

import (
	"fmt"
	"sort"
)

func Allergies(score uint) (allergies []string) {
	scoreToAllergen := map[int]string{
		128: "cats",
		64:  "pollen",
		32:  "chocolate",
		16:  "tomatoes",
		8:   "strawberries",
		4:   "shellfish",
		2:   "peanuts",
		1:   "eggs",
	}

	keys := make([]int, 0)
	for k := range scoreToAllergen {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	sort.Slice(keys, func(a, b int) bool {
		return keys[b] < keys[a]
	})

	for _, k := range keys {
		v := scoreToAllergen[k]
		fmt.Printf("%d/uint(%d)=%d\n", score, k, score/uint(k))
		if score/uint(k) == 1 {
			allergies = append(allergies, v)
			fmt.Printf("setting score to %v\n", score%uint(k))

			score = score % uint(k)
		}
	}
	return allergies
}

func AllergicTo(score uint, allergen string) bool {
	allergies := Allergies(score)
	fmt.Printf("score(%d) is allergic to %v\n", score, allergies)
	for _, candidate := range allergies {
		if candidate == allergen {
			return true
		}
	}
	return false
}
