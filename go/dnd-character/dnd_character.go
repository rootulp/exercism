package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	roll1 := rollDice()
	roll2 := rollDice()
	roll3 := rollDice()
	roll4 := rollDice()

	return getSum(roll1, roll2, roll3, roll4) - getMin(roll1, roll2, roll3, roll4)
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constituion := Ability()

	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constituion,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constituion),
	}
}

// rollDice returns a random number between 1 and 6.
func rollDice() int {
	return rand.Intn(5) + 1
}

func getSum(slice ...int) (sum int) {
	for _, x := range slice {
		sum += x
	}
	return sum
}

func getMin(slice ...int) (minimum int) {
	minimum = slice[0]
	for _, x := range slice {
		if x < minimum {
			minimum = x
		}
	}
	return minimum
}
