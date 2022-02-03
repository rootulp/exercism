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

var symbolToPlant = map[string]string{
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
	trimmed := strings.Trim(diagram, "\n")
	rows := strings.Split(trimmed, "\n")
	sort.Strings(children)
	return &Garden{diagram: rows, children: children}, nil
}

func (g *Garden) Plants(child string) (plants []string, ok bool) {
	fmt.Printf("diagrams %v\n", g.diagram)
	index := indexOf(g.children, child)
	column := index * 2
	cups := []string{
		string(g.diagram[0][column]),
		string(g.diagram[0][column+1]),
		string(g.diagram[1][column]),
		string(g.diagram[1][column+1]),
	}
	for _, cup := range cups {
		plants = append(plants, symbolToPlant[cup])
	}
	return plants, true
}

func indexOf(slice []string, element string) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}
