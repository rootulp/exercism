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

var verseNumberToMiddleLines = map[int][]string{
	2: {
		"It wriggled and jiggled and tickled inside her.",
		"She swallowed the spider to catch the fly.",
	},
	3: {
		"How absurd to swallow a bird!",
		"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
	},
	4: {
		"Imagine that, to swallow a cat!",
		"She swallowed the cat to catch the bird.",
	},
	5: {
		"What a hog, to swallow a dog!",
		"She swallowed the dog to catch the cat.",
	},
	6: {
		"Just opened her throat and swallowed a goat!",
		"She swallowed the goat to catch the dog.",
	},
	7: {
		"I don't know how she swallowed a cow!",
		"She swallowed the cow to catch the goat.",
	},
}

func Verse(verseNumber int) string {
	lines := []string{}
	lines = append(lines, firstLine(verseNumber))
	lines = append(lines, middleLines(verseNumber)...)
	lines = append(lines, lastLine(verseNumber))
	return strings.Join(lines, "\n")
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

func firstLine(verseNumber int) string {
	animal := startingAnimal(verseNumber)
	return fmt.Sprintf("I know an old lady who swallowed a %v.", animal)
}

func startingAnimal(verseNumber int) string {
	return verseNumberToStartingAnimal[verseNumber]
}

func middleLines(verseNumber int) []string {
	if verseNumber <= 1 {
		return []string{}
	}
	if verseNumber == 2 {
		return verseNumberToMiddleLines[verseNumber]
	}
	if verseNumber >= 3 && verseNumber <= 7 {
		lines := verseNumberToMiddleLines[verseNumber]
		return append(lines, middleLines(verseNumber - 1)[1:]...)
	}
	if verseNumber == 8 {
		return []string{}
	}
	panic(fmt.Sprintf("unsupported verseNumber %v", verseNumber))
}

func lastLine(verseNumber int) string {
	if verseNumber == 8 {
		return "She's dead, of course!"
	}
	return "I don't know why she swallowed the fly. Perhaps she'll die."
}
