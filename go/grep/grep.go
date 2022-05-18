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
	filename string
}

type configuration struct {
	// -n Print the line numbers of each matching line.
	prefixLineNumbers bool
	// -l Print only the names of files that contain at least one matching line.
	printFileNames bool
	// -i Match line using a case-insensitive comparison.
	matchCaseInsensitive bool
	// -v Invert the program -- collect all lines that fail to match the pattern.
	invertMatch bool
	// -x Only match entire lines, instead of lines that contain a match.
	matchEntireLine bool

	prefixFileName bool
}

func NewConfiguration(flags []string, files []string) (c configuration) {
	c = configuration{
		prefixLineNumbers:    false,
		printFileNames:       false,
		matchCaseInsensitive: false,
		invertMatch:          false,
		matchEntireLine:      false,
		prefixFileName:       false,
	}

	for _, flag := range flags {
		switch flag {
		case "-n":
			c.prefixLineNumbers = true
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

	if len(files) > 1 {
		c.prefixFileName = true
	}

	return c
}

func Search(pattern string, flags, files []string) (result []string) {
	lines := readFiles(files)
	config := NewConfiguration(flags, files)
	matches := search(pattern, config, lines)

	temp := format(matches, config)
	fmt.Printf("result %v\n", temp)
	return temp
}

func search(pattern string, config configuration, lines []line) (matches []line) {
	for _, line := range lines {
		if config.matchCaseInsensitive {
			lowerContents := strings.ToLower(line.contents)
			lowerPattern := strings.ToLower(pattern)
			if strings.Contains(lowerContents, lowerPattern) {
				matches = append(matches, line)
			}
		} else if config.matchEntireLine && config.invertMatch {
			if pattern != line.contents {
				matches = append(matches, line)
			}
		} else if config.matchEntireLine {
			if pattern == line.contents {
				matches = append(matches, line)
			}
		} else if config.invertMatch {
			if !strings.Contains(line.contents, pattern) {
				matches = append(matches, line)
			}
		} else {
			if strings.Contains(line.contents, pattern) {
				matches = append(matches, line)
			}
		}
	}
	return matches
}

func format(matches []line, config configuration) (result []string) {
	result = []string{}
	for _, match := range matches {
		if config.printFileNames {
			result = append(result, match.filename)
		} else if config.prefixFileName && config.prefixLineNumbers {
			result = append(result, fmt.Sprintf("%s:%d:%s", match.filename, match.number, match.contents))
		} else if config.prefixFileName {
			result = append(result, fmt.Sprintf("%s:%s", match.filename, match.contents))
		} else if config.prefixLineNumbers {
			result = append(result, fmt.Sprintf("%d:%s", match.number, match.contents))
		} else {
			result = append(result, match.contents)
		}
	}

	if config.printFileNames {
		return uniq(result)
	}
	return result
}

func readFiles(files []string) (lines []line) {
	for _, file := range files {
		lines = append(lines, readLines(file)...)
	}
	return lines
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
		line := line{lineNumber, contents, filename}

		lines = append(lines, line)
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func uniq(items []string) (result []string) {
	result = []string{}
	for _, item := range items {
		if !contains(result, item) {
			result = append(result, item)
		}
	}
	return result
}

func contains(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
