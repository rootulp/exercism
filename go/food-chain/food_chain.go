package foodchain

import (
	"fmt"
	"strings"
)

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
		return []string{"It wriggled and jiggled and tickled inside her.", "She swallowed the spider to catch the fly."}
	}
	if verseNumber == 3 {
		return []string{"How absurd to swallow a bird!", "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.", "She swallowed the spider to catch the fly."}
	}
	panic(fmt.Sprintf("unsupported verseNumber %v", verseNumber))
}

func startingAnimal(verseNumber int) string {
	switch verseNumber {
	case 1:
		return "fly"
	case 2:
		return "spider"
	case 3:
		return "bird"
	default:
		panic(fmt.Sprintf("unsupported verseNumber %v", verseNumber))
	}
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
