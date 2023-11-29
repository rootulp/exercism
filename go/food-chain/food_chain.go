package foodchain

import (
	"fmt"
	"strings"
)

func Verse(v int) string {
	if v == 1 {
		lines := []string{firstLine("fly"), lastLine()}
		return strings.Join(lines, "\n")
	}
	return "I know an old lady who swallowed a spider.\nIt wriggled and jiggled and tickled inside her.\nShe swallowed the spider to catch the fly.\nI don't know why she swallowed the fly. Perhaps she'll die."
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
