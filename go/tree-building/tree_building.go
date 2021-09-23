package tree

import "reflect"

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

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	nodes := map[int]*Node{}
	nodes[0] = &Node{ID: 0}

	for _, record := range records {
		// Initialize node if it doesn't exist
		if _, ok := nodes[record.ID]; !ok {
			nodes[record.ID] = &Node{ID: record.ID}
		}
		// Initialize parent if it doesn't exist
		if _, ok := nodes[record.Parent]; !ok {
			nodes[record.Parent] = &Node{ID: record.Parent}
		}
		// Populate children for parent
		children := nodes[record.Parent].Children
		node := nodes[record.ID]
		if !contains(children, node) {
			children = append(children, node)
			nodes[record.Parent].Children = children
		}
	}
	return nodes[0], nil
}

func contains(slice []*Node, node *Node) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, node) {
			return true
		}
	}
	return false
}
