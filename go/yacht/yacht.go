package yacht

import (
	"log"
	"reflect"
	"sort"
)

func Score(dice []int, category string) (score int) {
	switch category {
	case "ones":
		return 1 * countOf(dice, 1)
	case "twos":
		return 2 * countOf(dice, 2)
	case "threes":
		return 3 * countOf(dice, 3)
	case "fours":
		return 4 * countOf(dice, 4)
	case "fives":
		return 5 * countOf(dice, 5)
	case "sixes":
		return 6 * countOf(dice, 6)
	case "full house":
		return scoreFullHouse(dice)
	case "four of a kind":
		return scoreFourOfKind(dice)
	case "little straight":
		return scoreLittleStraight(dice)
	case "big straight":
		return scoreBigStraight(dice)
	case "choice":
		return sum(dice)
	case "yacht":
		return scoreYacht(dice)
	default:
		log.Fatalf("unrecognized category %v", category)
		return 0
	}
}

func scoreFullHouse(dice []int) (score int) {
	if !isFullHouse(dice) {
		return 0
	}
	return sum(dice)
}

func scoreFourOfKind(dice []int) (score int) {
	if !isFourOfKind(dice) {
		return 0
	}
	diceToOccurences := getOccurences(dice)
	for d, occurences := range diceToOccurences {
		if occurences >= 4 {
			return d * 4
		}
	}
	log.Fatalf("diceToOccurences did not contain dice with 4 occurences %v", diceToOccurences)
	return 0
}

func scoreLittleStraight(dice []int) (score int) {
	if !isLittleStraight(dice) {
		return 0
	}
	return 30
}

func scoreBigStraight(dice []int) (score int) {
	if !isBigStraight(dice) {
		return 0
	}
	return 30
}

func scoreYacht(dice []int) (score int) {
	if !isYacht(dice) {
		return 0
	}
	return 50
}

func countOf(dice []int, item int) (count int) {
	for _, d := range dice {
		if d == item {
			count++
		}
	}
	return count
}

func isFullHouse(dice []int) bool {
	diceToOccurences := getOccurences(dice)
	for _, occurences := range diceToOccurences {
		if occurences != 2 && occurences != 3 {
			return false
		}
	}
	return true
}

func isFourOfKind(dice []int) bool {
	diceToOccurences := getOccurences(dice)
	for _, occurences := range diceToOccurences {
		if occurences >= 4 {
			return true
		}
	}
	return false
}

func isLittleStraight(dice []int) bool {
	littleStraight := []int{1, 2, 3, 4, 5}
	sort.Ints(dice)
	return reflect.DeepEqual(dice, littleStraight)
}

func isBigStraight(dice []int) bool {
	bigStraight := []int{2, 3, 4, 5, 6}
	sort.Ints(dice)
	return reflect.DeepEqual(dice, bigStraight)
}

func isYacht(dice []int) bool {
	diceToOccurences := getOccurences(dice)
	for _, occurences := range diceToOccurences {
		if occurences != 5 {
			return false
		}
	}
	return true
}

func getOccurences(dice []int) (diceToOccurences map[int]int) {
	diceToOccurences = map[int]int{}
	for _, d := range dice {
		if _, ok := diceToOccurences[d]; !ok {
			diceToOccurences[d] = 0
		}
		diceToOccurences[d] += 1
	}
	return diceToOccurences
}

func sum(dice []int) (result int) {
	for _, d := range dice {
		result += d
	}
	return result
}
