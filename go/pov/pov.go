package pov

type Tree struct {
	value string
	children []*Tree
	parent *Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	node := &Tree{
		value: value,
		children: children,
	}
	for _, child := range children {
		child.parent = node
	}

	return node
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	fromTree := tr.findNode(from)
	if fromTree == nil {
		return nil
	}

	// constuct path from `from` to `root` to not lose track of parents during reversal
	path := []*Tree{fromTree}
	for p := fromTree.parent; p != nil; p = p.parent {
		path = append(path, p)
	}
	// from is already root
	if len(path) == 0 {
		return fromTree
	}

	fromTree.parent = nil
	for i := 0; i < len(path)-1; i++ {
		flip(path[i], path[i+1])
	}
	return fromTree
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	newPov := tr.FromPov(to)
	if newPov == nil {
		return nil
	}
	fromNode := newPov.findNode(from)
	if fromNode == nil {
		return nil
	}
	path := []string{fromNode.value}
	for p := fromNode.parent; p != nil; p = p.parent {
		path = append(path, p.value)
		if p.value == to {
			return path
		}
	}
	return nil
}

func (tr *Tree) findNode(value string) *Tree {
	if tr.value == value {
		return tr
	}
	for _, child := range tr.children {
		if found := child.findNode(value); found != nil {
			return found
		}
	}
	return nil
}

func flip(child *Tree, parent *Tree) {
	parent.parent = child
	parent.removeChild(child)
	child.addChild(parent)
}

func (tr *Tree) removeChild(child *Tree) {
	for i, ch := range tr.children {
		if ch == child {
			tr.children = append(tr.children[:i], tr.children[i+1:]...)
			return
		}
	}
}

func (tr *Tree) addChild(child *Tree) {
	tr.children = append(tr.children, child)
}
