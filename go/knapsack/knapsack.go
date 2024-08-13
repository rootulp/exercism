package knapsack

import (
	"fmt"
	"sort"
)

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) (maxValue int) {
	sorted := sortItemsByValuePerWeight(items)
	fmt.Printf("sorted: %v\n", sorted)
	currentWeight := maximumWeight
	for _, item := range sorted {
		if item.Weight <= currentWeight {
			currentWeight -= item.Weight
			maxValue += item.Value
		}
	}
	return maxValue
}

func sortItemsByValuePerWeight(items []Item) []Item {
	sort.Slice(items, func(i, j int) bool {
		iValuePerWeight := float64(items[i].Value) / float64(items[i].Weight)
		jValuePerWeight := float64(items[j].Value) / float64(items[j].Weight)
		return iValuePerWeight > jValuePerWeight
	})
	return items
}
