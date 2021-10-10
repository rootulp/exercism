package tournament

import (
	"fmt"
	"io"
	"strings"
)

type stat struct {
	MatchesPlayed int
	Wins          int
	Draws         int
	Loses         int
	Points        int
}

// Define a function Tally(io.Reader, io.Writer) error.
func Tally(r io.Reader, w io.Writer) error {
	lines, err := getLines(r)
	if err != nil {
		return err
	}
	teamsToScores := map[string]*stat{}
	for _, line := range lines {
		if isComment(line) {
			continue
		}
		// fmt.Printf("line: %s", line)
		parts := strings.Split(line, ";")
		if len(parts) < 3 {
			continue
		}
		fmt.Printf("parts: %s", parts)
		a := parts[0]
		b := parts[1]
		// _ := parts[2]

		if _, ok := teamsToScores[a]; !ok {
			teamsToScores[a] = &stat{}
		}
		if _, ok := teamsToScores[b]; !ok {
			teamsToScores[b] = &stat{}
		}

		teamsToScores[a].MatchesPlayed += 1
		teamsToScores[b].MatchesPlayed += 1
	}

	io.WriteString(w, getTable(teamsToScores))
	return nil
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#")
}

func getLines(r io.Reader) ([]string, error) {
	input, err := io.ReadAll(r)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(input), "\n")
	return lines, nil
}

func getTable(teamsToScores map[string]*stat) (result string) {
	// const header = "Team                           | MP |  W |  D |  L |  P"
	var header = fmt.Sprintf("%-30s |%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	result += header
	for team, stat := range teamsToScores {
		result += fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n", team, stat.MatchesPlayed, stat.Wins, stat.Draws, stat.Loses, stat.Points)
	}
	return result
}
