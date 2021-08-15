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
	stringA := getString(a)
	stringB := getString(b)
	return strings.Contains(stringB, stringA)
}

func getString(a []int) string {
	result := []string{}
	for _, val := range a {
		str := strconv.Itoa(val)
		result = append(result, str)
	}
	return strings.Join(result, ",")
}
