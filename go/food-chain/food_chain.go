package foodchain

import (
	"fmt"
	"strings"
)

var verseNumberToStartingAnimal = map[int]string{
	1: "fly",
	2: "spider",
	3: "bird",
	4: "cat",
	5: "dog",
	6: "goat",
	7: "cow",
	8: "horse",
}

func Verse(verseNumber int) string {
	lines := []string{firstLine(startingAnimal(verseNumber))}
	lines = append(lines, middleLines(verseNumber)...)
	lines = append(lines, lastLine())
	return strings.Join(lines, "\n")
}

func middleLines(verseNumber int) []string {
	if verseNumber <= 1 {
		return []string{}
	}
	if verseNumber == 2 {
		return []string{
			"It wriggled and jiggled and tickled inside her.",
			"She swallowed the spider to catch the fly.",
		}
	}
	if verseNumber == 3 {
		return []string{
			"How absurd to swallow a bird!",
			"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
			"She swallowed the spider to catch the fly.",
		}
	}
	if verseNumber == 4 {
		return []string{
			"Imagine that, to swallow a cat!",
			"She swallowed the cat to catch the bird.",
			"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
			"She swallowed the spider to catch the fly.",
		}
	}
	if verseNumber == 5 {
		return []string{
			"What a hog, to swallow a dog!",
			"She swallowed the dog to catch the cat.",
			"She swallowed the cat to catch the bird.",
			"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
			"She swallowed the spider to catch the fly.",
		}
	}
	panic(fmt.Sprintf("unsupported verseNumber %v", verseNumber))
}

func startingAnimal(verseNumber int) string {
	return verseNumberToStartingAnimal[verseNumber]
}

func firstLine(animal string) string {
	return fmt.Sprintf("I know an old lady who swallowed a %v.", animal)
}

func lastLine() string {
	return "I don't know why she swallowed the fly. Perhaps she'll die."
}

func Verses(start, end int) string {
	panic("Please implement the Verses function")
}

func Song() string {
	panic("Please implement the Song function")
}
