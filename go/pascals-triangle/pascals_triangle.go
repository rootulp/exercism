package pascal

func Triangle(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	previousRows := Triangle(n - 1)
	lastRow := previousRows[len(previousRows)-1]

	currentRow := []int{}
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			currentRow = append(currentRow, 1)
		} else {
			sum := lastRow[i-1] + lastRow[i]
			currentRow = append(currentRow, sum)
		}
	}
	return append(previousRows, currentRow)
}
