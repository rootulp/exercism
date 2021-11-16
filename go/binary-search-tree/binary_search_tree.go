package binarysearchtree

import "strconv"

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
	return SearchTreeData{data: i}
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
	if i <= std.data && std.left == nil {
		std.left = &SearchTreeData{data: i}
	} else if i > std.data && std.right == nil {
		std.right = &SearchTreeData{data: i}
	} else if i <= std.data && std.left != nil {
		std.left.Insert(i)
	} else if i > std.data && std.right != nil {
		std.right.Insert(i)
	}
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(func(int) string) (result []string) {
	for _, v := range visitInt(std) {
		result = append(result, strconv.Itoa(v))
	}
	return result
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(func(int) int) (result []int) {
	return visitInt(std)
}

func visitInt(std *SearchTreeData) (result []int) {
	if std == nil {
		return []int{}
	}
	result = append(result, visitInt(std.left)...)
	result = append(result, std.data)
	result = append(result, visitInt(std.right)...)
	return result
}
