package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) (filtered IntList) {
	filtered = make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (s IntList) Length() (length int) {
	// Assume we can't call len(s)
	for range s {
		length++
	}
	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	panic("Please implement the Map function")
}

func (s IntList) Reverse() IntList {
	panic("Please implement the Reverse function")
}

func (s IntList) Append(lst IntList) IntList {
	panic("Please implement the Append function")
}

func (s IntList) Concat(lists []IntList) IntList {
	panic("Please implement the Concat function")
}
