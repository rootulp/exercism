package knapsack

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	if len(items) == 0 {
		return 0
	}
	currentWeight := maximumWeight
	total := 0
	for _, item := range items {
		if item.Weight <= currentWeight {
			currentWeight -= item.Weight
			total += item.Value
		}
	}
	return total
}
