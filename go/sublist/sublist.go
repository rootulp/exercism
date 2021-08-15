package sublist

import (
	"reflect"
	"strconv"
	"strings"
)

type Relation string

func Sublist(a []int, b []int) Relation {
	if reflect.DeepEqual(a, b) {
		return "equal"
	} else if isSublist(a, b) {
		return "sublist"
	} else if isSublist(b, a) {
		return "superlist"
	} else {
		return "unequal"
	}
}

// isSublist returns true if a is a sublist of b
func isSublist(a []int, b []int) bool {
	return strings.Contains(getString(b), getString(a))
}

// getString returns a string representation of the slice provided
func getString(slice []int) string {
	strs := []string{}
	for _, val := range slice {
		str := strconv.Itoa(val)
		strs = append(strs, str)
	}
	return strings.Join(strs, ",")
}
