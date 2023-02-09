package change

import "fmt"

func Change(coins []int, target int) ([]int, error) {
	for _, coin := range coins {
		if coin == target {
			return []int{coin}, nil
		}
	}
	return []int{}, fmt.Errorf("no coins make %d change", target)
}
