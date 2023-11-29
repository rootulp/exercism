package foodchain

import (
	"fmt"
	"strings"
)

func startingAnimal(verseNumber int) string {
	switch verseNumber {
	case 1:
		return "fly"
	case 2:
		return "spider"
	default:
		panic(fmt.Sprintf("unsupported verseNumber %v", verseNumber))
	}
}

func Verse(verseNumber int) string {
	lines := []string{firstLine(startingAnimal(verseNumber))}
	lines = append(lines, middleLines(verseNumber)...)
	lines = append(lines, lastLine())
	return strings.Join(lines, "\n")
}

func middleLines(v int) []string {
	if v == 0 || v == 1 {
		return []string{}
	}
	if v == 2 {
		return []string{"It wriggled and jiggled and tickled inside her.", "She swallowed the spider to catch the fly."}
	}
	return []string{}
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
