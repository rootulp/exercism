package series

func All(n int, s string) (result []string) {
	for i := 0; i+n <= len(s); i += 1 {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	panic("Please implement the UnsafeFirst function")
}
