package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	diagram  []string
	children []string
}

var cupCodeToPlant = map[string]string{
	"R": "radishes",
	"C": "clover",
	"G": "grass",
	"V": "violets",
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !isValidDiagram(diagram) {
		return &Garden{}, fmt.Errorf("invalid diagram")
	}
	if !isValidNames(children) {
		return &Garden{}, fmt.Errorf("invalid names")
	}
	clone := make([]string, len(children))
	copy(clone, children)
	sort.Strings(clone)
	return &Garden{diagram: getDiagramRows(diagram), children: clone}, nil
}

func (g *Garden) Plants(child string) (plants []string, ok bool) {
	index := indexOf(g.children, child)
	if index == -1 {
		return []string{}, false
	}
	column := index * 2
	cups := []string{
		string(g.diagram[0][column]),
		string(g.diagram[0][column+1]),
		string(g.diagram[1][column]),
		string(g.diagram[1][column+1]),
	}
	for _, cup := range cups {
		plants = append(plants, cupCodeToPlant[cup])
	}
	return plants, true
}

func isValidNames(names []string) bool {
	seen := map[string]bool{}
	for _, name := range names {
		if _, ok := seen[name]; ok {
			// encountered duplicate name
			return false
		}
		seen[name] = true
	}
	return true
}

func isValidDiagram(diagram string) bool {
	return strings.HasPrefix(diagram, "\n") &&
		isEvenRows(diagram) &&
		isEvenCups(diagram) &&
		isValidCupCodes(diagram)
}

func isEvenRows(diagram string) bool {
	rows := getDiagramRows(diagram)
	return len(rows[0]) == len(rows[1])
}

func isEvenCups(diagram string) bool {
	rows := getDiagramRows(diagram)
	return len(rows[0])%2 == 0
}

func isValidCupCodes(diagram string) bool {
	rows := getDiagramRows(diagram)
	for _, row := range rows {
		for _, cupCode := range row {
			if _, ok := cupCodeToPlant[string(cupCode)]; !ok {
				return false
			}
		}
	}
	return true
}

func getDiagramRows(diagram string) (rows []string) {
	trimmed := strings.Trim(diagram, "\n")
	return strings.Split(trimmed, "\n")
}

func indexOf(slice []string, element string) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}
