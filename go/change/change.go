package change

import (
	"errors"
	"math"
)

// Change returns the minimum coins from available that add up to the target. It
// returns an error if target is negative or if not enough coins are available
// to make the target change.
func Change(coins []int, target int) ([]int, error) {
	changeToCoins := generateChangeToCoins(coins, target)
	result, ok := changeToCoins[target]
	if !ok {
		return nil, errors.New("no change available")
	}
	return result, nil
}

func generateChangeToCoins(coins []int, target int) map[int][]int {
	changeToCoins := map[int][]int{0: {}}
	for i := 1; i <= target; i++ {
		var min = math.MaxUint8
		for _, coin := range coins {
			remainingChange, ok := changeToCoins[i-coin]
			if ok && min > len(remainingChange) {
				min = len(remainingChange)
				changeToCoins[i] = append([]int{coin}, remainingChange...)
			}
		}
	}
	return changeToCoins
}
