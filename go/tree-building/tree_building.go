package tree

import (
	"errors"
	"sort"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// ByID implements sort.Interface based on the ID field.
type ByID []*Node

func (c ByID) Len() int {
	return len(c)
}

func (c ByID) Swap(a int, b int) {
	c[a], c[b] = c[b], c[a]
}

func (c ByID) Less(a int, b int) bool {
	return c[a].ID < c[b].ID
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodes := map[int]*Node{}
	nodes[0] = &Node{ID: 0}

	if rootNodeHasParent(records) {
		return &Node{}, errors.New("error root node has a parent")
	}
	if noRootNode(records) {
		return &Node{}, errors.New("error no root node")
	}
	if duplicateNode(records) {
		return &Node{}, errors.New("error duplicate node")
	}
	if nonContinuous(records) {
		return &Node{}, errors.New("error non-continuous node")
	}
	if directCycle(records) {
		return &Node{}, errors.New("error direct cycle")
	}

	for _, record := range records {
		// Initialize node if it doesn't exist
		if _, ok := nodes[record.ID]; !ok {
			nodes[record.ID] = &Node{ID: record.ID}
		}
		// Initialize parent if it doesn't exist
		if _, ok := nodes[record.Parent]; !ok {
			nodes[record.Parent] = &Node{ID: record.Parent}
		}
		// Continue if this is the root node
		if record.ID == 0 {
			continue
		}
		// Populate children for parent
		children := nodes[record.Parent].Children
		node := nodes[record.ID]
		if !contains(children, node) {
			children = append(children, node)
			sort.Sort(ByID(children))
			nodes[record.Parent].Children = children
		}
	}
	return nodes[0], nil
}

func rootNodeHasParent(records []Record) bool {
	for _, record := range records {
		if record.ID == 0 && record.Parent != 0 {
			return true
		}
	}
	return false
}

func noRootNode(records []Record) bool {
	for _, record := range records {
		if record.ID == 0 {
			return false
		}
	}
	return true
}

func duplicateNode(records []Record) bool {
	seen := map[int]bool{}
	for _, record := range records {
		if _, hasSeen := seen[record.ID]; hasSeen {
			return true
		}
		seen[record.ID] = true
	}
	return false
}

func nonContinuous(records []Record) bool {
	length := len(records)
	for _, record := range records {
		if record.ID > length-1 {
			return true
		}
	}
	return false
}

func directCycle(records []Record) bool {
	for _, record := range records {
		if record.ID == record.Parent && record.ID != 0 {
			return true
		}
	}
	return false
}

func contains(slice []*Node, node *Node) bool {
	for _, v := range slice {
		if v.ID == node.ID {
			return true
		}
	}
	return false
}
