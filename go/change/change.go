package change

import (
	"fmt"
	"math"
)

// Change returns the minimum coins from availableCoins that add up to the
// target. It returns an error if target is negative or if not enough coins are
// available to make the target change.
func Change(availableCoins []int, target int) ([]int, error) {
	sumToCoins := generateSumToCoins(availableCoins, target)
	result, ok := sumToCoins[target]
	if !ok {
		return nil, fmt.Errorf("no change available for target %d", target)
	}
	return result, nil
}

// generateSumToCoins returns a map from sum to the fewest coins that add up
// to that sum. The map spans the range from 0 to target inclusive.
func generateSumToCoins(availableCoins []int, target int) map[int][]int {
	sumToCoins := map[int][]int{0: {}}
	for i := 1; i <= target; i++ {
		var minNumCoins = math.MaxUint8
		for _, coin := range availableCoins {
			remainingChange, ok := sumToCoins[i-coin]
			if ok && minNumCoins > len(remainingChange) {
				minNumCoins = len(remainingChange)
				sumToCoins[i] = append([]int{coin}, remainingChange...)
			}
		}
	}
	return sumToCoins
}
