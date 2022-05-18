package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type line struct {
	number   int
	contents string
}

type configuration struct {
	// -n Print the line numbers of each matching line.
	printLineNumbers bool
	// -l Print only the names of files that contain at least one matching line.
	printFileNames bool
	// -i Match line using a case-insensitive comparison.
	matchCaseInsensitive bool
	// -v Invert the program -- collect all lines that fail to match the pattern.
	invertMatch bool
	// -x Only match entire lines, instead of lines that contain a match.
	matchEntireLine bool
}

func NewConfiguration(flags []string) (c configuration) {
	c = configuration{
		printLineNumbers:     false,
		printFileNames:       false,
		matchCaseInsensitive: false,
		invertMatch:          false,
		matchEntireLine:      false,
	}

	for _, flag := range flags {
		switch flag {
		case "-n":
			c.printLineNumbers = true
		case "-l":
			c.printFileNames = true
		case "-i":
			c.matchCaseInsensitive = true
		case "-v":
			c.invertMatch = true
		case "-x":
			c.matchEntireLine = true
		default:
			log.Fatalf("unrecognized flag %v", flag)
		}
	}

	return c
}

func Search(pattern string, flags, files []string) (result []string) {
	matches := []line{}
	config := NewConfiguration(flags)
	for _, file := range files {
		lines := readLines(file)
		for _, line := range lines {
			if strings.Contains(line.contents, pattern) {
				matches = append(matches, line)
			}
		}
	}
	return format(matches, config)
}

func format(matches []line, config configuration) (result []string) {
	for _, match := range matches {
		if config.printLineNumbers {
			result = append(result, fmt.Sprintf("%d:%s", match.number, match.contents))
		} else {
			result = append(result, match.contents)
		}
	}
	return result
}

func readLines(filename string) (lines []line) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		contents := scanner.Text()
		line := line{lineNumber, contents}

		lines = append(lines, line)
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
