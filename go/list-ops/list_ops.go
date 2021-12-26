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

func (s IntList) Map(fn func(int) int) (mapped IntList) {
	mapped = make(IntList, len(s))
	for i, v := range s {
		mapped[i] = fn(v)
	}
	return mapped
}

func (s IntList) Reverse() (reversed IntList) {
	reversed = make(IntList, 0)
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return reversed
}

func (s IntList) Append(toAppend IntList) (appended IntList) {
	return append(s, toAppend...)
}

func (s IntList) Concat(lists []IntList) (concatenated IntList) {
	concatenated = make(IntList, 0)
	concatenated = append(concatenated, s...)
	for _, list := range lists {
		concatenated = append(concatenated, list...)
	}
	return concatenated
}
