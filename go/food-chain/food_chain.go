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
		lines := []string{
			"How absurd to swallow a bird!",
			"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
		}
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 4 {
		lines := []string{
			"Imagine that, to swallow a cat!",
			"She swallowed the cat to catch the bird.",
		}
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 5 {
		lines := []string{
			"What a hog, to swallow a dog!",
			"She swallowed the dog to catch the cat.",
		}
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 6 {
		lines := []string{
			"Just opened her throat and swallowed a goat!",
			"She swallowed the goat to catch the dog.",
		}
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 7 {
		lines := []string{
			"I don't know how she swallowed a cow!",
			"She swallowed the cow to catch the goat.",
		}
		return append(lines, middleLines(verseNumber - 1)[1:]...)
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
