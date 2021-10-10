package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type stat struct {
	Name          string
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
		if isComment(line) || isEmptyLine(line) {
			continue
		}
		a, b, outcome, err := parseLine(line)
		if err != nil {
			return err
		}

		teamsToScores = initializeTeamIfNotExists(teamsToScores, a)
		teamsToScores = initializeTeamIfNotExists(teamsToScores, b)

		incrementMatchesPlayed(teamsToScores, a)
		incrementMatchesPlayed(teamsToScores, b)

		switch outcome {
		case "win":
			teamsToScores[a].Wins += 1
			teamsToScores[a].Points += 3
			teamsToScores[b].Loses += 1
		case "loss":
			teamsToScores[a].Loses += 1
			teamsToScores[b].Wins += 1
			teamsToScores[b].Points += 3
		case "draw":
			teamsToScores[a].Draws += 1
			teamsToScores[b].Draws += 1
			teamsToScores[a].Points += 1
			teamsToScores[b].Points += 1
		default:
			return fmt.Errorf("unexpected outcome %s", outcome)
		}
	}

	io.WriteString(w, getTable(teamsToScores))
	return nil
}

func getTable(teamsToScores map[string]*stat) (result string) {
	sorted := sortTeamsByPointsThenAlphabetically(teamsToScores)
	var header = fmt.Sprintf("%-30s |%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	result += header
	for _, stat := range sorted {
		result += fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n", stat.Name, stat.MatchesPlayed, stat.Wins, stat.Draws, stat.Loses, stat.Points)
	}
	return result
}

func sortTeamsByPointsThenAlphabetically(teamsToScores map[string]*stat) (result []*stat) {
	for _, stat := range teamsToScores {
		result = append(result, stat)
	}
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].Points != result[j].Points {
			return result[i].Points > result[j].Points
		}
		return result[i].Name < result[j].Name
	})

	return result
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#")
}

func isEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

func getLines(r io.Reader) ([]string, error) {
	input, err := io.ReadAll(r)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(input), "\n")
	return lines, nil
}

func incrementMatchesPlayed(teamsToScores map[string]*stat, team string) {
	teamsToScores[team].MatchesPlayed += 1
}

func initializeTeamIfNotExists(teamsToScores map[string]*stat, team string) map[string]*stat {
	if _, ok := teamsToScores[team]; !ok {
		teamsToScores[team] = &stat{
			Name: team,
		}
	}
	return teamsToScores
}

func parseLine(line string) (a string, b string, outcome string, err error) {
	parts := strings.Split(line, ";")
	if len(parts) < 3 {
		return "", "", "", fmt.Errorf("unexpected number of parts in line. Expected 3 parts delimited by semicolon. Got %d parts", len(parts))
	}
	return parts[0], parts[1], parts[2], nil
}
