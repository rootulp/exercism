package stringset

import (
	"fmt"
	"reflect"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	set map[string]bool
}

func New() Set {
	return Set{map[string]bool{}}
}

func NewFromSlice(slice []string) (result Set) {
	result = Set{map[string]bool{}}
	for _, element := range slice {
		result.Add(element)
	}
	return result
}

func (s Set) String() string {
	if s.set == nil {
		return "{}"
	}
	return fmt.Sprintf("{%s}", strings.Join(s.formattedKeys(), ", "))
}

func (s Set) IsEmpty() bool {
	return len(s.set) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.set[elem]
	return ok
}

func (s Set) Add(elem string) {
	s.set[elem] = true
}

func Subset(s1, s2 Set) bool {
	for key, val := range s1.set {
		if val && !s2.Has(key) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for key, val := range s1.set {
		if val && s2.Has(key) {
			return false
		}
	}
	for key, val := range s2.set {
		if val && s1.Has(key) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1.set, s2.set)
}

func Intersection(s1, s2 Set) (result Set) {
	result = New()
	for key, val := range s1.set {
		if val && s2.Has(key) {
			result.Add(key)
		}
	}
	return result
}

func Difference(s1, s2 Set) (result Set) {
	result = New()
	for key, val := range s1.set {
		if val && !s2.Has(key) {
			result.Add(key)
		}
	}
	return result
}

func Union(s1, s2 Set) (result Set) {
	result = New()
	for key := range s1.set {
		result.Add(key)
	}
	for key := range s2.set {
		result.Add(key)
	}
	return result
}

func (s Set) formattedKeys() (formatted []string) {
	formatted = []string{}
	for _, key := range s.keys() {
		quotted := fmt.Sprintf("\"%s\"", key)
		formatted = append(formatted, quotted)
	}
	return formatted
}

func (s Set) keys() (keys []string) {
	for key := range s.set {
		keys = append(keys, key)
	}
	return keys
}
