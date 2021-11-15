package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}
	result := []int{}
	for _, v := range i {
		if filter(v) {
			result = append(result, v)
		}
	}
	i = result
	return i
}

func (i Ints) Discard(filter func(int) bool) Ints {
	if i == nil {
		return nil
	}
	return i.Keep(func(i int) bool {
		return !filter(i)
	})
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return nil
	}
	panic("Please implement the Keep function")
}

func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return nil
	}
	panic("Please implement the Keep function")
}
