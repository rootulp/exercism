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
	lines = append(lines, lastLine(verseNumber))
	return strings.Join(lines, "\n")
}

func middleLines(verseNumber int) []string {
	if verseNumber <= 1 {
		return []string{}
	}
	if verseNumber == 2 {
		return verseNumberToMiddleLines[verseNumber]
	}
	if verseNumber == 3 {
		lines := verseNumberToMiddleLines[verseNumber]
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 4 {
		lines := verseNumberToMiddleLines[verseNumber]
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 5 {
		lines := verseNumberToMiddleLines[verseNumber]
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 6 {
		lines := verseNumberToMiddleLines[verseNumber]
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 7 {
		lines := verseNumberToMiddleLines[verseNumber]
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 8 {
		return []string{}
	}
	panic(fmt.Sprintf("unsupported verseNumber %v", verseNumber))
}

var verseNumberToMiddleLines = map[int][]string{
	2: {
		"It wriggled and jiggled and tickled inside her.",
		"She swallowed the spider to catch the fly.",
	},
	3: {
		"How absurd to swallow a bird!",
		"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
	},
	4: []string{
		"Imagine that, to swallow a cat!",
		"She swallowed the cat to catch the bird.",
	},
	5: []string{
		"What a hog, to swallow a dog!",
		"She swallowed the dog to catch the cat.",
	},
	6: []string{
		"Just opened her throat and swallowed a goat!",
		"She swallowed the goat to catch the dog.",
	},
	7: []string{
		"I don't know how she swallowed a cow!",
		"She swallowed the cow to catch the goat.",
	},
}

func startingAnimal(verseNumber int) string {
	return verseNumberToStartingAnimal[verseNumber]
}

func firstLine(animal string) string {
	return fmt.Sprintf("I know an old lady who swallowed a %v.", animal)
}

func lastLine(verseNumber int) string {
	if verseNumber == 8 {
		return "She's dead, of course!"
	}
	return "I don't know why she swallowed the fly. Perhaps she'll die."
}

func Verses(start, end int) string {
	verses := []string{}
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
