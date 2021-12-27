package series

func All(n int, s string) (result []string) {
	for i := 0; i+n <= len(s); i += 1 {
		result = append(result, s[i:i+n])
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
