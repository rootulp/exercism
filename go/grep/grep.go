package grep

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) (matches []string) {
	for _, file := range files {
		lines := readLines(file)
		for _, line := range lines {
			if strings.Contains(line, pattern) {
				matches = append(matches, line)
			}
		}
	}
	return matches
}

func readLines(filename string) (lines []string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
